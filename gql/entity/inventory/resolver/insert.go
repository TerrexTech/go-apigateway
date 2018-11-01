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
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_INVENTORY")

	// Marshal Inventory-data
	inventoryJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "InventoryInsertResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "InventoryInsertResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV1()
	if err != nil {
		err = errors.Wrap(err, "InventoryInsertResolver: Error generating UUID for InsertInventory-Event")
		return nil, err
	}

	// Publish Insert-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		Action:        "insert",
		CorrelationID: cid,
		AggregateID:   2,
		Data:          inventoryJSON,
		Timestamp:     time.Now(),
		TimeUUID:      eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, cid)
	if err != nil {
		err = errors.Wrap(err, "InventoryInsertResolver: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		insertInvResp := handleInsertInvResponse(kr)
		if insertInvResp != nil {
			if insertInvResp.Err == nil {
				return insertInvResp.Result, nil
			}
			insertInvErr := insertInvResp.Err
			err = errors.Wrap(insertInvErr.Err, "InventoryInsertResolver: InsertInventory Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: InsertInventory Error", insertInvErr.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleInsertInvResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(
			errors.New(kr.Error),
			"InventoryInsertResponseHandler: Error in KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	inventory := map[string]interface{}{}
	err := json.Unmarshal(kr.Result, &inventory)
	if err != nil {
		err = errors.Wrap(
			err,
			"InventoryInsertResponseHandler: "+
				"Error while Unmarshalling KafkaResponse into inventory",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: inventory,
		Err:    nil,
	}
}
