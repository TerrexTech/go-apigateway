package resolver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/gql/response"
	"github.com/TerrexTech/go-apigateway/gwerrors"

	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// genericQuery is a generic-resolver for Inventory GraphQL-query.
// Other queries call this function.
var genericQuery = func(serviceAction string, params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_WARNING")

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	paramsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "WarningQueryResolver: Error marshalling params into JSON")
		return nil, err
	}

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "WarningQueryResolver: Error generating UUID for cid")
		return nil, err
	}
	uuid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "WarningQueryResolver: Error generating UUID for Query-Event")
		return nil, err
	}

	// Publish Auth-Request on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "query",
		CorrelationID: cid,
		AggregateID:   18,
		Data:          paramsJSON,
		NanoTime:      time.Now().UnixNano(),
		ServiceAction: serviceAction,
		UUID:          uuid,
		YearBucket:    2018,
	}

	cio, err := kf.EnsureConsumerIO(consTopic, consTopic, false, uuid)
	if err != nil {
		err = errors.Wrap(err, "WarningQueryResolver: Error creating ConsumerIO")
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
			authRes := handleWarnQueryResponse(msg)
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

func handleWarnQueryResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "WarningQueryResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	result := []interface{}{}
	err := json.Unmarshal(kr.Result, &result)
	if err != nil {
		err = errors.Wrap(
			err,
			"WarningQueryResponseHandler: Error while Unmarshalling warning into KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	m := []map[string]interface{}{}
	for i, r := range result {
		item, assertOK := r.(map[string]interface{})
		if !assertOK {
			err = errors.New("error asserting item to map[string]interface{}")
			err = errors.Wrapf(
				err,
				"WarningQueryResponseHandler: Error asserting item at index: %d", i,
			)
			log.Println(err)
			krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
			return &response.ResolverResponse{
				Result: nil,
				Err:    krerr,
			}
		}

		m = append(m, item)
	}

	return &response.ResolverResponse{
		Result: m,
		Err:    nil,
	}
}
