package resolver

import (
	"os"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

var LoginResolver = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_LOGIN")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_LOGIN")

	rootValue := params.Info.RootValue.(map[string]interface{})
	ka := rootValue["kafkaAdapter"].(*kafka.Adapter)
	pio, err := ka.EnsureProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for LoginResolver")
		return nil, err
	}
	cio, err := ka.EnsureConsumerIO(consTopic, consTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for LoginResolver")
		return nil, err
	}
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	return authHandler(ts, params.Args, pio, cio)
}
