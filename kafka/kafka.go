package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/consumer"
	"github.com/TerrexTech/go-kafkautils/producer"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/pkg/errors"
)

// Adapter allows conveniently connecting to Adapter, and creates required
// Topics and channels for Adapter-communication.
type Adapter struct {
	Brokers           []string
	ConsumerGroupName string
	ConsumerTopics    []string
	ProducerTopic     string
}

// responseProducer creates a new Kafka-Producer used for producing the
// responses after processing consumed Kafka-messages.
func (ka *Adapter) responseProducer(
	brokers []string,
) (*producer.Producer, error) {
	config := producer.Config{
		KafkaBrokers: brokers,
	}
	resProducer, err := producer.New(&config)
	if err != nil {
		return nil, err
	}
	return resProducer, nil
}

// Consumer creates a new Kafka-Consumer which listens for the events.
func (ka *Adapter) consumer() (*consumer.Consumer, error) {
	saramaCfg := cluster.NewConfig()
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaCfg.Consumer.MaxProcessingTime = 10 * time.Second
	saramaCfg.Consumer.Return.Errors = true

	config := &consumer.Config{
		ConsumerGroup: ka.ConsumerGroupName,
		KafkaBrokers:  ka.Brokers,
		SaramaConfig:  saramaCfg,
		Topics:        ka.ConsumerTopics,
	}

	return consumer.New(config)
}

// InitIO initializes KafkaIO from the configuration provided to KafkaAdapter.
// It is necessary that both consumer and producer are properly setup, to enable
// response for every request. Else this operation will be marked as failed,
// and the service won't run.
func (ka *Adapter) InitIO() (*IO, error) {
	log.Println("Initializing KafkaIO")

	// Create Kafka Response-Producer
	resProducer, err := ka.responseProducer(ka.Brokers)
	if err != nil {
		err = errors.Wrap(err, "Error Creating Response Producer")
		return nil, err
	}
	log.Println("Create Kafka Response-Producer")
	resProducerInput, err := resProducer.Input()
	if err != nil {
		err = errors.Wrap(err, "Error Getting Input-Channel from Producer")
		return nil, err
	}

	// Setup Producer I/O channels
	producerInputChan := make(chan *model.KafkaResponse)
	kio := &IO{
		producerInputChan: (chan<- *model.KafkaResponse)(producerInputChan),
		producerErrChan:   resProducer.Errors(),
	}

	// The Kafka-Response post-processing the consumed events
	go func() {
		for msg := range producerInputChan {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				err = errors.Wrap(err, "Error Marshalling KafkaResponse for Aggregate")
				log.Println(err)
				return
			}

			producerMsg := producer.CreateMessage(ka.ProducerTopic, msgJSON)
			resProducerInput <- producerMsg
		}
	}()
	log.Println("Created Kafka Response-Channel")

	// Create Kafka Event-Consumer
	eventConsumer, err := ka.consumer()
	if err != nil {
		err = errors.Wrap(err, "Error Creating ConsumerGroup for Events")
		return nil, err
	}
	log.Println("Created Kafka Event-Consumer Group")

	// A channel which receives consumer-messages to be committed
	consumerOffsetChan := make(chan *sarama.ConsumerMessage)
	kio.consumerOffsetChan = (chan<- *sarama.ConsumerMessage)(consumerOffsetChan)
	go func() {
		for msg := range consumerOffsetChan {
			eventConsumer.MarkOffset(msg, "")
		}
	}()
	log.Println("Created Kafka Event Offset-Commit Channel")

	// Setup Consumer I/O channels
	kio.consumerErrChan = eventConsumer.Errors()
	kio.consumerMsgChan = eventConsumer.Messages()
	log.Println("KafkaIO Ready")

	return kio, nil
}
