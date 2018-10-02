package util

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/producer"
	"github.com/pkg/errors"
)

// producer creates a new Kafka-Producer used for producing the
// responses after processing consumed Kafka-messages.
func (ka *KafkaFactory) producer(
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

func (ka *KafkaFactory) newProducerIO(id string) (*producerIO, error) {
	// Create Kafka Response-Producer
	resProducer, err := ka.producer(ka.Brokers)
	if err != nil {
		err = errors.Wrap(err, "Error Creating Kafka-Producer")
		return nil, err
	}
	log.Println("Creating Kafka-Producer")
	resProducerInput, err := resProducer.Input()
	if err != nil {
		err = errors.Wrap(err, "Error getting Input-Channel from Kafka-Producer")
		return nil, err
	}

	// Setup Producer I/O channels
	producerInputChan := make(chan *sarama.ProducerMessage)
	pio := &producerIO{
		pid:               id,
		producerInputChan: (chan<- *sarama.ProducerMessage)(producerInputChan),
		producerErrChan:   resProducer.Errors(),
	}

	// The Kafka-Response post-processing the consumed events
	go func() {
		for msg := range producerInputChan {
			resProducerInput <- msg
		}
	}()
	log.Println("ProducerIO Ready")
	return pio, nil
}

func (ka *KafkaFactory) newEventProducerIO(
	id string, enableErrors bool,
) (*EventProducerIO, error) {
	pio, err := ka.newProducerIO(id)
	if err != nil {
		err = errors.Wrap(err, "Error creating Event-ProducerIO")
	}

	inputChan := make(chan *model.Event)
	epio := &EventProducerIO{
		id:        id,
		inputChan: (chan<- *model.Event)(inputChan),
		errChan:   pio.errors(),
	}

	if !enableErrors {
		go func() {
			for err := range epio.errChan {
				parsedErr := errors.Wrap(err.Err, "Event-Producer Error")
				log.Println(parsedErr)
				log.Println(err)
			}
		}()
	}

	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_REGISTER")
	go func() {
		for msg := range inputChan {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				err = errors.Wrapf(err,
					"Error Marshalling KafkaResponse with CorrelationID: %s and AggregateID: %d, "+
						"on topic %s",
					msg.CorrelationID,
					msg.AggregateID,
					prodTopic,
				)
				log.Println(err)
				continue
			}

			producerMsg := producer.CreateMessage(prodTopic, msgJSON)
			pio.input() <- producerMsg
		}
	}()
	log.Println("KafkaResponse-ProducerIO Ready")
	return epio, nil
}

func (ka *KafkaFactory) newKafkaResponseProducerIO(
	id string, enableErrors bool,
) (*KafkaResponseProducerIO, error) {
	pio, err := ka.newProducerIO(id)
	if err != nil {
		err = errors.Wrap(err, "Error creating KafkaResponse-ProducerIO")
	}

	inputChan := make(chan *model.KafkaResponse)
	krpio := &KafkaResponseProducerIO{
		id:        id,
		inputChan: (chan<- *model.KafkaResponse)(inputChan),
		errChan:   pio.errors(),
	}

	if !enableErrors {
		go func() {
			for err := range krpio.errChan {
				parsedErr := errors.Wrap(err.Err, "KafkaResponse-Producer Error")
				log.Println(parsedErr)
				log.Println(err)
			}
		}()
	}

	// The Kafka-Response post-processing the consumed events
	go func() {
		for msg := range inputChan {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				err = errors.Wrapf(err,
					"Error Marshalling KafkaResponse with CorrelationID: %s and AggregateID: %d, "+
						"on topic %s",
					msg.CorrelationID,
					msg.AggregateID,
					msg.Topic,
				)
				log.Println(err)
				continue
			}

			producerMsg := producer.CreateMessage(msg.Topic, msgJSON)
			pio.input() <- producerMsg
		}
	}()
	log.Println("KafkaResponse-ProducerIO Ready")
	return krpio, nil
}

func (ka *KafkaFactory) EnsureEventProducerIO(
	topic string,
	id string,
	enableErrors bool,
) (*EventProducerIO, error) {
	if epioStore[id] == nil {
		p, err := ka.newEventProducerIO(id, enableErrors)
		if err != nil {
			err = errors.Wrap(err, "Error creating ProducerEventIO")
			return nil, err
		}
		epioStore[id] = p
	}
	return epioStore[id], nil
}

func (ka *KafkaFactory) EnsureKafkaResponseProducerIO(
	id string,
	enableErrors bool,
) (*KafkaResponseProducerIO, error) {
	if krpioStore[id] == nil {
		p, err := ka.newKafkaResponseProducerIO(id, enableErrors)
		if err != nil {
			err = errors.Wrap(err, "Error creating ProducerIO")
			return nil, err
		}
		krpioStore[id] = p
	}
	return krpioStore[id], nil
}
