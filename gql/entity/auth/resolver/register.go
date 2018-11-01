package resolver

import (
	"context"
	"encoding/json"
	"fmt"
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

// Register is the GraphQL resolver for Register-endpoint.
var Register = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_USERAUTH")

	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error generating UUID for cid")
		return nil, err
	}
	timeUUID, err := uuuid.NewV1()
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error generating UUID for Register-Event")
		return nil, err
	}

	// Publish Auth-Request on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		Action:        "insert",
		CorrelationID: cid,
		AggregateID:   1,
		Data:          credentialsJSON,
		Timestamp:     time.Now(),
		TimeUUID:      timeUUID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, timeUUID)
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		registerResp := handleRegisterResponse(kr, ts)
		if registerResp != nil {
			if registerResp.Err == nil {
				return registerResp.Result, nil
			}
			ae := registerResp.Err
			err = errors.Wrap(ae.Err, "AuthRegisterResolver: Registration Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: Registration Error", ae.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleRegisterResponse(
	kr esmodel.KafkaResponse,
	ts auth.TokenStoreI,
) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(errors.New(kr.Error), "AuthRegisterResponseHandler: Error in KafkaResponse")
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
			"AuthRegisterResponseHandler: Error while Unmarshalling user into KafkaResponse",
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
			"AuthRegisterResponseHandler: Error generating AccessToken",
		)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResponseHandler: Error generating RefreshToken")
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
