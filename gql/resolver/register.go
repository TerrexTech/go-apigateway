package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/gwerrors"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-apigateway/model"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

type registerResponse struct {
	authResponse *model.AuthResponse
	authErr      *gwerrors.KRError
}

var RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
	log.Println("00000000000000000000000000")
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_REGISTER")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_REGISTER")

	log.Println("111111111111111111111111")
	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}
	log.Println(">>>>>Credentials")
	log.Println(string(credentialsJSON))

	log.Println("22222222222222222222222222")
	rootValue := params.Info.RootValue.(map[string]interface{})
	ka := rootValue["kafkaAdapter"].(*kafka.Adapter)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	log.Println("333333333333333333333333333")
	pio, err := ka.EnsureProducerEventIO(prodTopic, prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for RegisterResolver")
		return nil, err
	}
	cio, err := ka.EnsureConsumerIO(consTopic, consTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for RegisterResolver")
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

	log.Println("444444444444444444444444")
	// Publish Auth-Request on Kafka Topic
	go func() {
		log.Println("55555555555555555555")
		pio.ProducerInput() <- &esmodel.Event{
			Action:        "insert",
			CorrelationID: cid,
			AggregateID:   1,
			Data:          credentialsJSON,
			Timestamp:     time.Now(),
			UUID:          eventID,
			YearBucket:    2018,
		}
	}()

	log.Println("666666666666666666666")
	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

registerResponseLoop:
	// Check login-response messages for matching CorrelationID and return result
	for {
		select {
		case <-ctx.Done():
			break registerResponseLoop
		case msg := <-cio.ConsumerMessages():
			log.Println("777777777777777777777")
			cio.MarkOffset() <- msg
			log.Println("8888888888888888888888888")
			loginRes := handleRegisterResponse(msg, ts, cid)
			if loginRes != nil {
				log.Println("------------------")
				log.Println(loginRes.authErr)
				log.Println(loginRes.authResponse)
				if loginRes.authErr == nil {
					log.Println("*******************((((((((")
					return loginRes.authResponse, nil
				}
				ae := loginRes.authErr
				log.Println(ae)
				outErr := fmt.Errorf("%d: Registeration Error", ae.Code)
				return nil, outErr
			}
		}
	}

	return nil, errors.New("Timed out")
}

func handleRegisterResponse(
	msg *sarama.ConsumerMessage,
	ts auth.TokenStoreI,
	cid uuuid.UUID,
) *registerResponse {
	log.Println("8888888888888888888888888")
	user, krerr := parseKafkaResponse(msg, cid)
	if krerr != nil {
		err := errors.Wrap(krerr.Err, "Error authenticating user")
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &registerResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}
	if user == nil && krerr == nil {
		return nil
	}
	log.Println("9999999999999999999999999")

	at, err := genAccessToken(user)
	if err != nil {
		err = errors.Wrap(err, "Error generating AccessToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &registerResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "Error generating RefreshToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &registerResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}

	return &registerResponse{
		authResponse: &model.AuthResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
		authErr: nil,
	}
}
