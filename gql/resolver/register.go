package resolver

import (
	"encoding/json"
	"os"
	"time"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/gocql/gocql"

	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

var RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("EVENT_PRODUCER_TOPIC")
	// consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_REGISTER")

	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	uuid, err := gocql.RandomUUID()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for event")
		return nil, err
	}
	event := &model.Event{
		Action:      "insert",
		AggregateID: 1,
		Data:        string(credentialsJSON),
		Timestamp:   time.Now(),
		UserID:      1,
		UUID:        uuid,
		Version:     0,
		YearBucket:  2018,
	}

	rootValue := params.Info.RootValue.(map[string]interface{})
	ka := rootValue["kafkaAdapter"].(*kafka.Adapter)
	pio, err := ka.EnsureProducerEventIO(prodTopic, prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for RegisterResolver")
		return nil, err
	}

	pio.ProducerInput() <- event
	return nil, nil
}
