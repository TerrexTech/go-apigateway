package resolver

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"

	"github.com/graphql-go/graphql"
)

var FlashSaleAdd = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_FLASHADD")
	consGroup := os.Getenv("KAFKA_CONSUMER_GROUP_FLASHADD")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_FLASHADD")

	var aggregateID int8 = 4

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	flashAddJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "FlashSaleAddResolver: Error marshalling flash sale - Add into JSON")
		return nil, err
	}

	eventProdIo, err := kf.EnsureEventProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for FlashSaleAddResolver")
		return nil, err
	}

	//CorrelationID
	correlationID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "FlashSaleAddResolver: Error generating UUID for cid")
		return nil, err
	}

	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "FlashSaleAddResolver: Error generating UUID for FlashSaleAdd-Event")
		return nil, err
	}

	//Publish request on kafka topic
	go func() {
		eventProdIo.Input() <- &esmodel.Event{
			Action:        "insert",
			CorrelationID: correlationID,
			AggregateID:   aggregateID,
			Data:          flashAddJSON,
			Timestamp:     time.Now(),
			UUID:          eventID,
			YearBucket:    2018,
			Version:       1,
		}
	}()

	//Timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	krChan, err := kf.EnsureConsumerIO(consGroup, consTopic, false, correlationID)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for FlashSaleAddResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		flashAddResponse := handleFlashAddResponse(kr, correlationID)
		if flashAddResponse != nil {
			return flashAddResponse, nil
		}
	}
	return nil, errors.New("Unknown error - kr-FlashAddResponse")
}

type FlashAddResp struct {
	FlashAdd model.Flash
}

func handleFlashAddResponse(kr esmodel.KafkaResponse, cid uuuid.UUID) *FlashAddResp {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "handleFlashAddResponse: Error in KafkaResponse")
	}

	flashAddRes := model.Flash{}
	err := json.Unmarshal(kr.Result, &flashAddRes)
	if err != nil {
		err = errors.Wrap(err, "handleFlashAddResponse: Error while unmarshalling flash-add into KafkaResponse")
	}

	return &FlashAddResp{flashAddRes}
}
