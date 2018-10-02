package resolver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gwerrors"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// LoginResolver is the resolver for Login GraphQL query.
var LoginResolver = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_LOGIN")
	consGroup := os.Getenv("KAFKA_CONSUMER_GROUP_LOGIN")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_LOGIN")

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error generating UUID for cid")
		return nil, err
	}
	krpio, err := kf.EnsureKafkaResponseProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for LoginResolver")
		return nil, err
	}
	// Publish Auth-Request on Kafka Topic
	go func() {
		krpio.Input() <- &esmodel.KafkaResponse{
			CorrelationID: cid,
			Input:         credentialsJSON,
			Topic:         prodTopic,
		}
	}()

	cio, err := kf.EnsureConsumerIO(consGroup, consTopic, false, cid)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for LoginResolver")
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
			authRes := handleLoginResponse(msg, ts, cid)
			if authRes != nil {
				if authRes.err == nil {
					return authRes.tokens, nil
				}
				return nil, errors.New(authRes.err.Error())
			}
		}
	}

	return nil, errors.New("Timed out")
}

func handleLoginResponse(
	kr esmodel.KafkaResponse,
	ts auth.TokenStoreI,
	cid uuuid.UUID,
) *authResponse {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "LoginResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &authResponse{
			tokens: nil,
			err:    krerr,
		}
	}

	user := &model.User{}
	err := json.Unmarshal([]byte(kr.Result), user)
	if err != nil {
		err = errors.Wrap(
			err,
			"LoginResponseHandler: Error while Unmarshalling user into KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			tokens: nil,
			err:    krerr,
		}
	}

	at, err := genAccessToken(user)
	if err != nil {
		err = errors.Wrap(
			err,
			"LoginResponseHandler: Error generating AccessToken",
		)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			tokens: nil,
			err:    krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error generating RefreshToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			tokens: nil,
			err:    krerr,
		}
	}

	return &authResponse{
		tokens: &model.AuthTokens{
			AccessToken:  at,
			RefreshToken: rt,
		},
		err: nil,
	}
}
