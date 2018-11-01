package util

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-kafkautils/kafka"
	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

// EnsureConsumerIO creates a new Kafka-Consumer if one does not exist for that specific
// ID. Otherwise an existing cached Consumer is returned.
// ID here refers to "GroupName + TopicName" for the Consumer.
// This is not analogous to traditional Consumers that simply forward any Kafka Messages.
// This returns a new channel for every request. The request's response is then sent on that
// channel when the response arrives. It determines the correct response for request by using
// the UUID. If a response arrives before the request for it is registered, it caches
// the response until the request for it is received, and the response is sent out as soon
// request arrives.
func (kf *KafkaFactory) EnsureConsumerIO(
	group string,
	topic string,
	enableErrors bool,
	uuid uuuid.UUID,
) (<-chan model.KafkaResponse, error) {
	consCache := kf.cacheConfig
	if consCache.topicMapLocks[topic] == nil {
		consCache.topicMapLocks[topic] = &sync.RWMutex{}
	}

	id := group + topic
	if consCache.cioStore[id] == nil {
		// Create Kafka Aggregate-Response Consumer
		resConsumer, err := newResponseConsumer(
			kf.ctx,
			kf.errGroup,
			&kafka.ConsumerConfig{
				GroupName:    group,
				KafkaBrokers: kf.Brokers,
				Topics:       []string{topic},
			},
		)
		if err != nil {
			err = errors.Wrap(err, "Error Creating Response ConsumerGroup for Events")
			return nil, err
		}
		log.Println("--> Created Kafka Response ConsumerGroup")

		consCache.cioStore[id] = resConsumer

		// Handle Messages
		kf.errGroup.Go(func() error {
			for {
				select {
				case <-kf.ctx.Done():
					log.Printf("--> Closing Consumer-loop on Topic: %s", topic)
					return errors.New("service-context closed")
				case msg := <-resConsumer.Messages():
					go handleKafkaConsumerMsg(
						msg,
						consCache.uuidMap,
						consCache.uuidMapLock,
						consCache.topicMapLocks[msg.Topic],
					)
				}
			}
		})
	}

	// Update UUIDs for that specific topic
	consCache.uuidMapLock.RLock()
	topicCIDMap := consCache.uuidMap[topic]
	consCache.uuidMapLock.RUnlock()

	if topicCIDMap == nil {
		consCache.uuidMapLock.Lock()
		consCache.uuidMap[topic] = map[string]chan model.KafkaResponse{}
		consCache.uuidMapLock.Unlock()

		consCache.uuidMapLock.RLock()
		topicCIDMap = consCache.uuidMap[topic]
		consCache.uuidMapLock.RUnlock()
	}

	tl := consCache.topicMapLocks[topic]

	tl.RLock()
	topicCID := topicCIDMap[uuid.String()]
	tl.RUnlock()
	if topicCID == nil {
		// 1 buffer because channel might be read later after its written
		readChan := make(chan model.KafkaResponse, 1)
		tl.Lock()
		topicCIDMap[uuid.String()] = readChan
		tl.Unlock()

		tl.RLock()
		topicCID = topicCIDMap[uuid.String()]
		tl.RUnlock()
	}

	return (<-chan model.KafkaResponse)(topicCID), nil
}

// handleKafkaConsumerMsg handles the messages from ConsumerIO. It checks if any
// corresponding UUID exists for that response, and then passes the response to
// the channel associated with that UUID.
// If no UUID exists at time of message-arrival, the message is cached until a
// request is received, and the message is then sent as soon as the request is received.
// Once the response is sent, the corresponding UUID-entry is deleted from map.
func handleKafkaConsumerMsg(
	msg *sarama.ConsumerMessage,
	uuidMap map[string]map[string]chan model.KafkaResponse,
	uuidMapLock *sync.RWMutex,
	topicMapLock *sync.RWMutex,
) {
	kr := model.KafkaResponse{}
	err := json.Unmarshal(msg.Value, &kr)
	if err != nil {
		err = errors.Wrap(err, "Error Unmarshalling Kafka-Message into Kafka-Response")
		log.Println(err)
		return
	}
	kruuid := kr.UUID.String()

	uuidMapLock.RLock()
	topicCIDMap := uuidMap[msg.Topic]
	uuidMapLock.RUnlock()

	// Create topic-entry in CID map since it doesn't exist
	if topicCIDMap == nil {
		uuidMapLock.Lock()
		uuidMap[msg.Topic] = make(map[string]chan model.KafkaResponse)
		uuidMapLock.Unlock()

		uuidMapLock.RLock()
		topicCIDMap = uuidMap[msg.Topic]
		uuidMapLock.RUnlock()
	}

	// Get the associated channel from TopicCIDMap
	topicMapLock.RLock()
	exists := topicCIDMap[kruuid] != nil
	topicMapLock.RUnlock()

	// Send the response to the channel from above and delete the entry from map
	if exists {
		topicMapLock.Lock()
		topicMapLock.Unlock()

		topicCIDMap[kruuid] <- kr
		delete(topicCIDMap, kruuid)
		return
	}

	// If no corresponding UUID was found, we cache the response until a CID
	// requesting that response is available.
	readChan := make(chan model.KafkaResponse, 1)
	topicMapLock.Lock()
	topicCIDMap[kruuid] = readChan
	topicMapLock.Unlock()
	readChan <- kr
}
