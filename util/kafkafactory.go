package util

import (
	"github.com/TerrexTech/go-kafkautils/consumer"
)

var (
	cidMap = make(map[string]map[string]CIDSubAdapter)

	epioStore  = map[string]*EventProducerIO{}
	krpioStore = map[string]*KafkaResponseProducerIO{}
	cioStore   = map[string]*consumer.Consumer{}
)

// KafkaFactory allows conveniently connecting to KafkaFactory, and creates required
// Topics and channels for KafkaFactory-communication.
type KafkaFactory struct {
	Brokers []string
}
