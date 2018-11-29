package resolver

import (
	"context"
	"encoding/json"
	"fmt"
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

// Insert is the GraphQL resolver for InsertInventory GraphQL query.
var Insert = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_WASTE")

	// Marshal Inventory-data
	disposalJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "DisposalInsertResponseHandler: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DisposalInsertResponseHandler: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DisposalInsertResponseHandler: Error generating UUID for InsertInventory-Event")
		return nil, err
	}

	// Publish Insert-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "insert",
		CorrelationID: cid,
		AggregateID:   10,
		Data:          disposalJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "DisposalInsertResponseHandler: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		insertDisposeResp := handleInsertDisposeResponse(kr)
		if insertDisposeResp != nil {
			if insertDisposeResp.Err == nil {
				return insertDisposeResp.Result, nil
			}
			insertDisposeErr := insertDisposeResp.Err
			err = errors.Wrap(insertDisposeErr.Err, "DisposalInsertResponseHandler: InsertDisposal Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: InsertDisposal Error", insertDisposeErr.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleInsertDisposeResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(
			errors.New(kr.Error),
			"DisposalInsertResponseHandler: Error in KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	disposal := map[string]interface{}{}
	err := json.Unmarshal(kr.Result, &disposal)
	if err != nil {
		err = errors.Wrap(
			err,
			"DisposalInsertResponseHandler: "+
				"Error while Unmarshalling KafkaResponse into disposal",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: disposal,
		Err:    nil,
	}
}
