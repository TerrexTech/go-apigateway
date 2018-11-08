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

type DeleteResult struct {
	DeletedCount int64 `json:"deletedCount"`
}

// Delete is the GraphQL resolver for DeleteInventory-endpoint.
var Delete = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_INVENTORY")

	// Marshal Inventory-data
	deleteJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "InventoryDeleteResolver: Error marshalling into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "InventoryDeleteResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "InventoryDeleteResolver: Error generating UUID for DeleteInventory-Event")
		return nil, err
	}

	// Publish Update-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "delete",
		CorrelationID: cid,
		AggregateID:   2,
		Data:          deleteJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "InventoryDeleteResolver: Error creating ConsumerIO for InventoryDeleteResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		DeleteInvResp := handleDeleteInvResponse(kr)
		if DeleteInvResp != nil {
			if DeleteInvResp.Err == nil {
				return DeleteInvResp.Result, nil
			}
			ae := DeleteInvResp.Err
			err = errors.Wrap(ae.Err, "InventoryDeleteResolver: DeleteInventory Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: DeleteInventory Error", ae.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleDeleteInvResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(errors.New(kr.Error), "DeleteInventoryResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	deleteResult := &DeleteResult{}
	err := json.Unmarshal(kr.Result, deleteResult)
	if err != nil {
		err = errors.Wrap(
			err,
			"DeleteInventoryResponseHandler: "+
				"Error while unmarshalling KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: deleteResult,
		Err:    nil,
	}
}
