package util

import (
	"github.com/TerrexTech/go-kafkautils/consumer"
)

var (
	// Event Producers cache
	epioStore = map[string]*EventProducerIO{}
	// Kafka-Response Producers cache
	krpioStore = map[string]*KafkaResponseProducerIO{}
	// Consumers Cache
	cioStore = map[string]*consumer.Consumer{}
)

// KafkaFactory allows conveniently creating required Kafka Producers and Consumers.
type KafkaFactory struct {
	Brokers []string
}
