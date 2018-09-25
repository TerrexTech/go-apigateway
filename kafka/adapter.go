package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/consumer"
	"github.com/TerrexTech/go-kafkautils/producer"
	"github.com/TerrexTech/uuuid"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/pkg/errors"
)

var groupIDSuffix = func() string {
	id, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "KafkaAdapter: Error generating GroupIDSuffix")
		log.Fatalln(err)
	}
	return id.String()
}()

var pioStore = map[string]*ProducerIO{}
var peioStore = map[string]*ProducerEventIO{}
var cioStore = map[string]*ConsumerIO{}

// Adapter allows conveniently connecting to Adapter, and creates required
// Topics and channels for Adapter-communication.
type Adapter struct {
	Brokers []string
}

// producer creates a new Kafka-Producer used for producing the
// responses after processing consumed Kafka-messages.
func (ka *Adapter) producer(
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
func (ka *Adapter) consumer(name string, topics []string) (*consumer.Consumer, error) {
	saramaCfg := cluster.NewConfig()
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaCfg.Consumer.MaxProcessingTime = 10 * time.Second
	saramaCfg.Consumer.Return.Errors = true

	config := &consumer.Config{
		ConsumerGroup: name,
		KafkaBrokers:  ka.Brokers,
		SaramaConfig:  saramaCfg,
		Topics:        topics,
	}

	return consumer.New(config)
}

func (ka *Adapter) newConsumerIO(
	name string,
	topics []string,
	enableErrors bool,
) (*ConsumerIO, error) {
	// Create Kafka Event-Consumer
	reqConsumer, err := ka.consumer(name, topics)
	if err != nil {
		err = errors.Wrap(err, "Error Creating Response ConsumerGroup for Events")
		return nil, err
	}
	log.Println("Created Kafka Response ConsumerGroup")

	// A channel which receives consumer-messages to be committed
	consumerOffsetChan := make(chan *sarama.ConsumerMessage)
	cio := &ConsumerIO{
		consumerErrChan:    reqConsumer.Errors(),
		consumerMsgChan:    reqConsumer.Messages(),
		consumerOffsetChan: (chan<- *sarama.ConsumerMessage)(consumerOffsetChan),
	}

	if !enableErrors {
		go func() {
			for err := range reqConsumer.Errors() {
				parsedErr := errors.Wrap(err, "Producer Error")
				log.Println(parsedErr)
				log.Println(err)
			}
		}()
	}

	go func() {
		for msg := range consumerOffsetChan {
			reqConsumer.MarkOffset(msg, "")
		}
	}()
	log.Println("Created Kafka Response Offset-Commit Channel")

	log.Println("ConsumerIO Ready")
	return cio, nil
}

func (ka *Adapter) newProducerIO(id string, enableErrors bool) (*ProducerIO, error) {
	// Create Kafka Response-Producer
	resProducer, err := ka.producer(ka.Brokers)
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
	pio := &ProducerIO{
		id:                id,
		producerInputChan: (chan<- *model.KafkaResponse)(producerInputChan),
		producerErrChan:   resProducer.Errors(),
	}

	if !enableErrors {
		go func() {
			for err := range resProducer.Errors() {
				parsedErr := errors.Wrap(err.Err, "Producer Error")
				log.Println(parsedErr)
				log.Println(err)
			}
		}()
	}

	// The Kafka-Response post-processing the consumed events
	go func() {
		for msg := range producerInputChan {
			// No need to unnecessarily increase message payload-size, so we remove "Topic".
			topic := msg.Topic
			msg.Topic = ""

			msgJSON, err := json.Marshal(msg)
			if err != nil {
				err = errors.Wrapf(err,
					"Error Marshalling KafkaResponse with CorrelationID: %s and AggregateID: %d, "+
						"on topic %s",
					msg.CorrelationID,
					msg.AggregateID,
					topic,
				)
				log.Println(err)
				return
			}

			producerMsg := producer.CreateMessage(topic, msgJSON)
			resProducerInput <- producerMsg
		}
	}()
	log.Println("ProducerIO Ready")
	return pio, nil
}

func (ka *Adapter) newProducerEventIO(topic string, id string, enableErrors bool) (*ProducerEventIO, error) {
	// Create Kafka Response-Producer
	resProducer, err := ka.producer(ka.Brokers)
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
	producerInputChan := make(chan *model.Event)
	pio := &ProducerEventIO{
		id:                id,
		producerInputChan: (chan<- *model.Event)(producerInputChan),
		producerErrChan:   resProducer.Errors(),
	}

	if !enableErrors {
		go func() {
			for err := range resProducer.Errors() {
				parsedErr := errors.Wrap(err.Err, "Producer Error")
				log.Println(parsedErr)
				log.Println(err)
			}
		}()
	}

	// The Kafka-Response post-processing the consumed events
	go func() {
		for msg := range producerInputChan {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				err = errors.Wrapf(err,
					"Error Marshalling Event with UUID: %s and AggregateID: %d, "+
						"on topic %s",
					msg.UUID,
					msg.AggregateID,
					topic,
				)
				log.Println(err)
				return
			}

			producerMsg := producer.CreateMessage(topic, msgJSON)
			resProducerInput <- producerMsg
		}
	}()
	log.Println("ProducerIO Ready")
	return pio, nil
}

func (ka *Adapter) EnsureProducerEventIO(
	topic string,
	id string,
	enableErrors bool,
) (*ProducerEventIO, error) {
	if peioStore[id] == nil {
		p, err := ka.newProducerEventIO(topic, id, enableErrors)
		if err != nil {
			err = errors.Wrap(err, "Error creating ProducerEventIO")
			return nil, err
		}
		peioStore[id] = p
	}
	return peioStore[id], nil
}

func (ka *Adapter) EnsureConsumerIO(
	id string,
	topic string,
	enableErrors bool,
) (*ConsumerIO, error) {
	if cioStore[id] == nil {
		name := id + "-" + groupIDSuffix
		c, err := ka.newConsumerIO(name, []string{topic}, enableErrors)
		if err != nil {
			err = errors.Wrap(err, "Error creating ConsumerIO")
			return nil, err
		}
		cioStore[id] = c
	}
	return cioStore[id], nil
}

func (ka *Adapter) EnsureProducerIO(
	id string,
	enableErrors bool,
) (*ProducerIO, error) {
	if pioStore[id] == nil {
		p, err := ka.newProducerIO(id, enableErrors)
		if err != nil {
			err = errors.Wrap(err, "Error creating ProducerIO")
			return nil, err
		}
		pioStore[id] = p
	}
	return pioStore[id], nil
}
