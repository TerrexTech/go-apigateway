package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
)

// IO provides channels for interacting with Kafka.
// Note: All receive-channels must be read from to prevent deadlock.
type IO struct {
	consumerErrChan    <-chan error
	consumerMsgChan    <-chan *sarama.ConsumerMessage
	consumerOffsetChan chan<- *sarama.ConsumerMessage
	producerErrChan    <-chan *sarama.ProducerError
	producerInputChan  chan<- *model.KafkaResponse
}

// ConsumerErrors returns send-channel where consumer errors are published.
func (kio *IO) ConsumerErrors() <-chan error {
	return kio.consumerErrChan
}

// ConsumerMessages returns send-channel where consumer messages are published.
func (kio *IO) ConsumerMessages() <-chan *sarama.ConsumerMessage {
	return kio.consumerMsgChan
}

// MarkOffset marks the consumer message-offset to be committed.
// This should be used once a message has done its job.
func (kio *IO) MarkOffset() chan<- *sarama.ConsumerMessage {
	return kio.consumerOffsetChan
}

// ProducerErrors returns send-channel where producer errors are published.
func (kio *IO) ProducerErrors() <-chan *sarama.ProducerError {
	return kio.producerErrChan
}

// ProducerInput returns receive-channel where kafka-responses can be produced.
func (kio *IO) ProducerInput() chan<- *model.KafkaResponse {
	return kio.producerInputChan
}
