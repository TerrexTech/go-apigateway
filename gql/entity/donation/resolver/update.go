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

type UpdateResult struct {
	MatchedCount  int64 `json:"matchedCount"`
	ModifiedCount int64 `json:"modifiedCount"`
}

// Update is the GraphQL resolver for UpdateInventory-endpoint.
var Update = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_DONATION")

	// Marshal Inventory-data
	updateJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "DonationUpdateResolver: Error marshalling into JSON")
		return nil, err
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DonationUpdateResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "DonationUpdateResolver: Error generating UUID for UpdateDonation-Event")
		return nil, err
	}

	// Publish Update-Event on Kafka Topic
	kf.EventProducer() <- &esmodel.Event{
		EventAction:   "update",
		CorrelationID: cid,
		AggregateID:   9,
		Data:          updateJSON,
		NanoTime:      time.Now().UnixNano(),
		UUID:          eventID,
		YearBucket:    2018,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consTopic, consTopic, false, eventID)
	if err != nil {
		err = errors.Wrap(err, "DonationUpdateResolver: Error creating ConsumerIO for DonationUpdateResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		UpdateDonateResp := handleUpdateDonateResponse(kr)
		if UpdateDonateResp != nil {
			if UpdateDonateResp.Err == nil {
				return UpdateDonateResp.Result, nil
			}
			ae := UpdateDonateResp.Err
			err = errors.Wrap(ae.Err, "DonationUpdateResolver: UpdateDonation Error")
			log.Println(err)
			outErr := fmt.Errorf("%d: UpdateDonation Error", ae.Code)
			return nil, outErr
		}
	}
	return nil, errors.New("Unknown Error")
}

func handleUpdateDonateResponse(kr esmodel.KafkaResponse) *response.ResolverResponse {
	if kr.Error != "" {
		err := errors.Wrap(errors.New(kr.Error), "UpdateDonationResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &response.ResolverResponse{
			Result: nil,
			Err:    krerr,
		}
	}

	updateResult := &UpdateResult{}
	err := json.Unmarshal(kr.Result, updateResult)
	if err != nil {
		err = errors.Wrap(
			err,
			"UpdateDonationResponseHandler: "+
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
		Result: updateResult,
		Err:    nil,
	}
}
