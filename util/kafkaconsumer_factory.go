package util

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/consumer"
	"github.com/TerrexTech/uuuid"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/pkg/errors"
)

// consumer creates a new Kafka-Consumer which listens for the events.
func (kf *KafkaFactory) consumer(
	name string, topics []string,
) (*consumer.Consumer, error) {
	saramaCfg := cluster.NewConfig()
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaCfg.Consumer.MaxProcessingTime = 10 * time.Second
	saramaCfg.Consumer.Return.Errors = true

	config := &consumer.Config{
		ConsumerGroup: name,
		KafkaBrokers:  kf.Brokers,
		SaramaConfig:  saramaCfg,
		Topics:        topics,
	}

	return consumer.New(config)
}

var (
	cidMap        = make(map[string]map[string]chan model.KafkaResponse)
	cidMapLock    = &sync.RWMutex{}
	topicMapLocks = make(map[string]*sync.RWMutex)
)

// EnsureConsumerIO creates a new Kafka-Consumer if one does not exist for that specific
// ID. Otherwise an existing cached Consumer is returned.
// ID here refers to "GroupName + TopicName" for the Consumer.
// This is not analogous to traditional Consumers that simply forward any Kafka Messages.
// This returns a new channel for every request. The request's response is then sent on that
// channel when the response arrives. It determines the correct response for request by using
// the CorrelationID. If a response arrives before the request for it is registered, it caches
// the response until the request for it is received, and the response is sent out as soon
// request arrives.
func (kf *KafkaFactory) EnsureConsumerIO(
	group string,
	topic string,
	enableErrors bool,
	cid uuuid.UUID,
) (<-chan model.KafkaResponse, error) {
	if topicMapLocks[topic] == nil {
		topicMapLocks[topic] = &sync.RWMutex{}
	}

	id := group + topic
	if cioStore[id] == nil {
		// Create Kafka Event-Consumer
		reqConsumer, err := kf.consumer(group, []string{topic})
		if err != nil {
			err = errors.Wrap(err, "Error Creating Response ConsumerGroup for Events")
			return nil, err
		}
		log.Println("Created Kafka Response ConsumerGroup")

		cioStore[id] = reqConsumer

		// Handle Messages
		go func() {
			for msg := range reqConsumer.Messages() {
				reqConsumer.MarkOffset(msg, "")
				go handleKafkaConsumerMsg(
					msg,
					cidMap,
					cidMapLock,
					topicMapLocks[msg.Topic],
				)
			}
		}()
	}

	// Update CorrelationIDs for that specific topic
	cidMapLock.RLock()
	topicCIDMap := cidMap[topic]
	cidMapLock.RUnlock()

	if topicCIDMap == nil {
		cidMapLock.Lock()
		cidMap[topic] = map[string]chan model.KafkaResponse{}
		cidMapLock.Unlock()

		cidMapLock.RLock()
		topicCIDMap = cidMap[topic]
		cidMapLock.RUnlock()
	}

	// 1 buffer because channel might be read later after its written
	readChan := make(chan model.KafkaResponse, 1)
	tl := topicMapLocks[topic]
	tl.Lock()
	topicCIDMap[cid.String()] = readChan
	tl.Unlock()

	return (<-chan model.KafkaResponse)(readChan), nil
}

// handleKafkaConsumerMsg handles the messages from ConsumerIO. It checks if any
// corresponding CorrelationID exists for that response, and then passes the response to
// the channel associated with that CorrelationID.
// If no CorrelationID exists at time of message-arrival, the message is cached until a
// request is received, and the message is then sent as soon as the request is received.
// Once the response is sent, the corresponding CorrelationID-entry is deleted from map.
func handleKafkaConsumerMsg(
	msg *sarama.ConsumerMessage,
	cidMap map[string]map[string]chan model.KafkaResponse,
	cidMapLock *sync.RWMutex,
	topicMapLock *sync.RWMutex,
) {
	kr := model.KafkaResponse{}
	err := json.Unmarshal(msg.Value, &kr)
	if err != nil {
		err = errors.Wrap(err, "Error Unmarshalling Kafka-Message into Kafka-Response")
		log.Println(err)
		return
	}
	krcid := kr.CorrelationID.String()

	cidMapLock.RLock()
	topicCIDMap := cidMap[msg.Topic]
	cidMapLock.RUnlock()

	// Create topic-entry in CID map since it doesn't exist
	if topicCIDMap == nil {
		cidMapLock.Lock()
		cidMap[msg.Topic] = make(map[string]chan model.KafkaResponse)
		cidMapLock.Unlock()

		cidMapLock.RLock()
		topicCIDMap = cidMap[msg.Topic]
		cidMapLock.RUnlock()
	}

	// Get the associated channel from TopicCIDMap
	topicMapLock.RLock()
	krReadChan, exists := topicCIDMap[krcid]
	topicMapLock.RUnlock()

	// Send the response to the channel from above and delete the entry from map
	if exists {
		topicMapLock.Lock()
		delete(topicCIDMap, krcid)
		topicMapLock.Unlock()

		krReadChan <- kr
		return
	}

	// If no corresponding CorrelationID was found, we cache the response until a CID
	// requesting that response is available.
	readChan := make(chan model.KafkaResponse, 1)
	topicMapLock.Lock()
	topicCIDMap[krcid] = readChan
	topicMapLock.Unlock()
}
