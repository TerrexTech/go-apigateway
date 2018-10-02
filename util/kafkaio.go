package util

import (
	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
)

type producerIO struct {
	pid               string
	producerErrChan   <-chan *sarama.ProducerError
	producerInputChan chan<- *sarama.ProducerMessage
}

func (pio *producerIO) id() string {
	return pio.pid
}

// ProducerErrors returns send-channel where producer errors are published.
func (pio *producerIO) errors() <-chan *sarama.ProducerError {
	return pio.producerErrChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (pio *producerIO) input() chan<- *sarama.ProducerMessage {
	return pio.producerInputChan
}

type KafkaResponseProducerIO struct {
	id        string
	errChan   <-chan *sarama.ProducerError
	inputChan chan<- *model.KafkaResponse
}

func (pio *KafkaResponseProducerIO) ID() string {
	return pio.id
}

// ProducerErrors returns send-channel where producer errors are published.
func (pio *KafkaResponseProducerIO) Errors() <-chan *sarama.ProducerError {
	return pio.errChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (pio *KafkaResponseProducerIO) Input() chan<- *model.KafkaResponse {
	return pio.inputChan
}

type EventProducerIO struct {
	id        string
	errChan   <-chan *sarama.ProducerError
	inputChan chan<- *model.Event
}

func (epio *EventProducerIO) ID() string {
	return epio.id
}

// ProducerErrors returns send-channel where producer errors are published.
func (epio *EventProducerIO) Errors() <-chan *sarama.ProducerError {
	return epio.errChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (epio *EventProducerIO) Input() chan<- *model.Event {
	return epio.inputChan
}
