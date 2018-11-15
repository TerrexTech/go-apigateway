package resolver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/entity/auth/model"
	"github.com/TerrexTech/go-apigateway/gql/response"
	"github.com/TerrexTech/go-apigateway/gwerrors"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// Login is the resolver for Login GraphQL-query.
var Login = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_USERAUTH")

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "AuthLoginResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "AuthLoginResolver: Error generating UUID for cid")
		return nil, err
	}
	uuid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "AuthLoginResolver: Error generating UUID for Login-Event")
		return nil, err
	}

	// Publish Auth-Request on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "query",
		CorrelationID: cid,
		AggregateID:   1,
		Data:          credentialsJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          uuid,
		YearBucket:    2018,
	}

	log.Println("%+v", &esmodel.Event{
		EventAction:   "query",
		CorrelationID: cid,
		AggregateID:   1,
		Data:          credentialsJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          uuid,
		YearBucket:    2018,
	})

	cio, err := kf.EnsureConsumerIO(consTopic, consTopic, false, uuid)
	if err != nil {
		err = errors.Wrap(err, "AuthLoginResolver: Error creating ConsumerIO")
		return nil, err
	}
	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

authResponseLoop:
	// Check auth-response messages for matching CorrelationID and return result
	for {
		select {
		case <-ctx.Done():
			break authResponseLoop
		case msg := <-cio:
			authRes := handleLoginResponse(msg, ts)
			if authRes != nil {
				if authRes.Err == nil {
					return authRes.Result, nil
				}
				return nil, errors.New(authRes.Err.Error())
			}
		}
	}

	return nil, errors.New("Timed out")
}

func handleLoginResponse(
	kr esmodel.KafkaResponse,
	ts auth.TokenStoreI,
) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "AuthLoginResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	user := &model.User{}
	err := json.Unmarshal(kr.Result, user)
	if err != nil {
		err = errors.Wrap(
			err,
			"AuthLoginResponseHandler: Error while Unmarshalling user into KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	at, err := genAccessToken(user)
	if err != nil {
		err = errors.Wrap(
			err,
			"AuthLoginResponseHandler: Error generating AccessToken",
		)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "AuthLoginResponseHandler: Error generating RefreshToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: &model.AuthTokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
		Err: nil,
	}
}
