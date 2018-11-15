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
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_SALE")

	// Marshal Sale-data
	saleJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "InvInsertResponseHandler: Error marshalling credentials into JSON")
		return nil, err
	}
	log.Println(string(saleJSON))

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "InvInsertResponseHandler: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "InvInsertResponseHandler: Error generating UUID for InsertSale-Event")
		return nil, err
	}

	// Publish Insert-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "insert",
		CorrelationID: cid,
		AggregateID:   3,
		Data:          saleJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "InvInsertResponseHandler: Error creating ConsumerIO")
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
			err = errors.Wrap(insertInvErr.Err, "InvInsertResponseHandler: InsertSale Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: InsertSale Error", insertInvErr.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleInsertInvResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(
			errors.New(kr.Error),
			"InvInsertResponseHandler: Error in KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	sale := map[string]interface{}{}
	err := json.Unmarshal(kr.Result, &sale)
	if err != nil {
		err = errors.Wrap(
			err,
			"InvInsertResponseHandler: "+
				"Error while Unmarshalling KafkaResponse into sale",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: sale,
		Err:    nil,
	}
}
