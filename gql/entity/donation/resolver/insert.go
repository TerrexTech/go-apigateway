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
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_DONATION")

	// Marshal Inventory-data
	donationJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "DonationInsertResponseHandler: Error marshalling credentials into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DonationInsertResponseHandler: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DonationInsertResponseHandler: Error generating UUID for InsertInventory-Event")
		return nil, err
	}

	// Publish Insert-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "insert",
		CorrelationID: cid,
		AggregateID:   9,
		Data:          donationJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "DonationInsertResponseHandler: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		insertDonateResp := handleInsertDonateResponse(kr)
		if insertDonateResp != nil {
			if insertDonateResp.Err == nil {
				return insertDonateResp.Result, nil
			}
			insertDonateErr := insertDonateResp.Err
			err = errors.Wrap(insertDonateErr.Err, "DonationInsertResponseHandler: InsertDonation Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: InsertDonation Error", insertDonateErr.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleInsertDonateResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(
			errors.New(kr.Error),
			"DonationInsertResponseHandler: Error in KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	donation := map[string]interface{}{}
	err := json.Unmarshal(kr.Result, &donation)
	if err != nil {
		err = errors.Wrap(
			err,
			"DonationInsertResponseHandler: "+
				"Error while Unmarshalling KafkaResponse into donation",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	return &response.ResolverResponse{
		Result: donation,
		Err:    nil,
	}
}
