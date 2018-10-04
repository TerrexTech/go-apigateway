package util

import (
	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
)

// producerIO creates a producer for generic Kafka-Messages.
type producerIO struct {
	producerErrChan   <-chan *sarama.ProducerError
	producerInputChan chan<- *sarama.ProducerMessage
}

// errors returns send-channel using which producer errors are published.
func (pio *producerIO) errors() <-chan *sarama.ProducerError {
	return pio.producerErrChan
}

// input returns receive-channel where kafka-responses can be produced.
func (pio *producerIO) input() chan<- *sarama.ProducerMessage {
	return pio.producerInputChan
}

// KafkaResponseProducerIO creates a producer for Kafka-Message of type KafkaResponse.
type KafkaResponseProducerIO struct {
	errChan   <-chan *sarama.ProducerError
	inputChan chan<- *model.KafkaResponse
}

// Errors returns send-channel where producer errors are published.
func (pio *KafkaResponseProducerIO) Errors() <-chan *sarama.ProducerError {
	return pio.errChan
}

// Input returns receive-channel using which KafkaResponses can be produced.
func (pio *KafkaResponseProducerIO) Input() chan<- *model.KafkaResponse {
	return pio.inputChan
}

// EventProducerIO creates a producer for Kafka-Message of type Events.
type EventProducerIO struct {
	errChan   <-chan *sarama.ProducerError
	inputChan chan<- *model.Event
}

// Errors returns send-channel where producer errors are published.
func (epio *EventProducerIO) Errors() <-chan *sarama.ProducerError {
	return epio.errChan
}

// Input returns receive-channel using which Events can be produced.
func (epio *EventProducerIO) Input() chan<- *model.Event {
	return epio.inputChan
}
