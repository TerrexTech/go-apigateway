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
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_WARNING")

	// Marshal Inventory-data
	deleteJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "WarningDeleteResolver: Error marshalling into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "WarningDeleteResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "WarningDeleteResolver: Error generating UUID for DeleteWarning-Event")
		return nil, err
	}

	// Publish Update-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "delete",
		CorrelationID: cid,
		AggregateID:   18,
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
		err = errors.Wrap(err, "WarningDeleteResolver: Error creating ConsumerIO for WarningDeleteResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		DeleteWarnResp := handleDeleteWarnResponse(kr)
		if DeleteWarnResp != nil {
			if DeleteWarnResp.Err == nil {
				return DeleteWarnResp.Result, nil
			}
			ae := DeleteWarnResp.Err
			err = errors.Wrap(ae.Err, "WarningDeleteResolver: DeleteWarning Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: DeleteWarning Error", ae.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleDeleteWarnResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(errors.New(kr.Error), "DeleteWarningResponseHandler: Error in KafkaResponse")
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
			"DeleteWarningResponseHandler: "+
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
