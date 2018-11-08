package util

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/kafka"
	"github.com/pkg/errors"
)

type KafkaConsCacheConfig struct {
	cioStore      map[string]*responseConsumer
	uuidMap       map[string]map[string]chan model.KafkaResponse
	uuidMapLock   *sync.RWMutex
	topicMapLocks map[string]*sync.RWMutex
}

// KafkaFactory allows conveniently creating required Kafka Producers and Consumers.
type KafkaFactory struct {
	Brokers  []string
	ctx      context.Context
	errGroup *errgroup.Group

	eventChan    chan *model.Event
	eventProduer *kafka.Producer
	cacheConfig  *KafkaConsCacheConfig
}

func NewKafkaFactory(
	ctx context.Context,
	errGroup *errgroup.Group,
	brokers []string,
	eventTopic string,
) (*KafkaFactory, error) {
	eventChan := make(chan *model.Event, 256)
	p, err := kafka.NewProducer(&kafka.ProducerConfig{
		KafkaBrokers: brokers,
	})
	if err != nil {
		err = errors.Wrap(err, "Error creating EventProducer")
		return nil, err
	}

	errGroup.Go(func() error {
		var prodErr error
	errLoop:
		for {
			select {
			case <-ctx.Done():
				break errLoop
			case err := <-p.Errors():
				if err != nil && err.Err != nil {
					parsedErr := errors.Wrap(err.Err, "Error in EventProducer")
					log.Println(parsedErr)
					log.Println(err)
					prodErr = err.Err
					break errLoop
				}
			}
		}
		log.Println("--> Closed EventProducer error-routine")
		return prodErr
	})

	closeProducer := false
	errGroup.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				closeProducer = true
				p.Close()
				log.Println("--> Closed EventProducer")
				return errors.New("EventProducer exited")
				// Replace provided version with MaxVersion from DB and send the query to
				// EventStoreQuery service.
			case event := <-eventChan:
				msgJSON, err := json.Marshal(event)
				if err != nil {
					err = errors.Wrapf(err,
						"Error Marshalling KafkaResponse with UUID: %s and AggregateID: %d, ",
						event.UUID,
						event.AggregateID,
					)
					log.Println(err)
				}

				if !closeProducer {
					p.Input() <- kafka.CreateMessage(eventTopic, msgJSON)
				} else {
					log.Println("--> Closed producer before producing Event")
				}
			}
		}
	})

	cacheConfig := &KafkaConsCacheConfig{
		cioStore:      map[string]*responseConsumer{},
		uuidMap:       map[string]map[string]chan model.KafkaResponse{},
		uuidMapLock:   &sync.RWMutex{},
		topicMapLocks: map[string]*sync.RWMutex{},
	}
	return &KafkaFactory{
		Brokers:  brokers,
		ctx:      ctx,
		errGroup: errGroup,

		eventChan:    eventChan,
		cacheConfig:  cacheConfig,
		eventProduer: p,
	}, nil
}

func (kf *KafkaFactory) EventProducer() chan<- *model.Event {
	return (chan<- *model.Event)(kf.eventChan)
}
