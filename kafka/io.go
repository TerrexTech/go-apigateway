package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
)

// IO provides channels for interacting with Kafka.
// Note: All receive-channels must be read from to prevent deadlock.
type ConsumerIO struct {
	consumerErrChan    <-chan error
	consumerMsgChan    <-chan *sarama.ConsumerMessage
	consumerOffsetChan chan<- *sarama.ConsumerMessage
}

// ConsumerErrors returns send-channel where consumer errors are published.
func (cio *ConsumerIO) ConsumerErrors() <-chan error {
	return cio.consumerErrChan
}

// ConsumerMessages returns send-channel where consumer messages are published.
func (cio *ConsumerIO) ConsumerMessages() <-chan *sarama.ConsumerMessage {
	return cio.consumerMsgChan
}

// MarkOffset marks the consumer message-offset to be committed.
// This should be used once a message has done its job.
func (cio *ConsumerIO) MarkOffset() chan<- *sarama.ConsumerMessage {
	return cio.consumerOffsetChan
}

type ProducerIO struct {
	id                string
	producerErrChan   <-chan *sarama.ProducerError
	producerInputChan chan<- *model.KafkaResponse
}

func (pio *ProducerIO) ID() string {
	return pio.id
}

// ProducerErrors returns send-channel where producer errors are published.
func (pio *ProducerIO) ProducerErrors() <-chan *sarama.ProducerError {
	return pio.producerErrChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (pio *ProducerIO) ProducerInput() chan<- *model.KafkaResponse {
	return pio.producerInputChan
}

type ProducerEventIO struct {
	id                string
	producerErrChan   <-chan *sarama.ProducerError
	producerInputChan chan<- *model.Event
}

func (pio *ProducerEventIO) ID() string {
	return pio.id
}

// ProducerErrors returns send-channel where producer errors are published.
func (pio *ProducerEventIO) ProducerErrors() <-chan *sarama.ProducerError {
	return pio.producerErrChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (pio *ProducerEventIO) ProducerInput() chan<- *model.Event {
	return pio.producerInputChan
}
