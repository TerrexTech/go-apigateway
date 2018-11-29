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

// Insert is the GraphQL resolver for InsertSale GraphQL query.
var Insert = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_FLASHSALE")

	// Marshal Inventory-data
	flashsaleJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "FlashsaleInsertResponseHandler: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "FlashsaleInsertResponseHandler: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "FlashsaleInsertResponseHandler: Error generating UUID for InsertInventory-Event")
		return nil, err
	}

	// Publish Insert-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "insert",
		CorrelationID: cid,
		AggregateID:   7,
		Data:          flashsaleJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "FlashsaleInsertResponseHandler: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		insertFlashResp := handleInsertFlashResponse(kr)
		if insertFlashResp != nil {
			if insertFlashResp.Err == nil {
				return insertFlashResp.Result, nil
			}
			insertFlashErr := insertFlashResp.Err
			err = errors.Wrap(insertFlashErr.Err, "FlashsaleInsertResponseHandler: InsertFlashsale Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: InsertFlashsale Error", insertFlashErr.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleInsertFlashResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(
			errors.New(kr.Error),
			"FlashsaleInsertResponseHandler: Error in KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	flashsale := map[string]interface{}{}
	err := json.Unmarshal(kr.Result, &flashsale)
	if err != nil {
		err = errors.Wrap(
			err,
			"FlashsaleInsertResponseHandler: "+
				"Error while Unmarshalling KafkaResponse into flashsale",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: flashsale,
		Err:    nil,
	}
}
