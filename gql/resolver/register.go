package resolver

import (
	"context"
	"encoding/json"
	"fmt"
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

// RegisterResolver is the GraphQL resolver for Register-endpoint.
var RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_REGISTER")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_REGISTER")

	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	ka := rootValue["kafkaFactory"].(*util.KafkaFactory)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	epio, err := ka.EnsureEventProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for RegisterResolver")
		return nil, err
	}
	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for Register-Event")
		return nil, err
	}

	// Publish Auth-Request on Kafka Topic
	go func() {
		epio.Input() <- &esmodel.Event{
			Action:        "insert",
			CorrelationID: cid,
			AggregateID:   1,
			Data:          credentialsJSON,
			Timestamp:     time.Now(),
			UUID:          eventID,
			YearBucket:    2018,
		}
	}()

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := ka.EnsureConsumerIO(consTopic, consTopic, false, cid)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for RegisterResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		registerResp := handleRegisterResponse(kr, ts)
		if registerResp != nil {
			if registerResp.err == nil {
				return registerResp.tokens, nil
			}
			ae := registerResp.err
			err = errors.Wrap(ae.Err, "Registration Error")
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
) *authResponse {
	if kr.Error != "" {
		err := errors.Wrap(errors.New(kr.Error), "RegisterResponseHandler: Error in KafkaResponse")
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
			"RegisterResponseHandler: Error while Unmarshalling user into KafkaResponse",
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
			"RegisterResponseHandler: Error generating AccessToken",
		)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			tokens: nil,
			err:    krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "Error generating RefreshToken")
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
