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

// Consumer creates a new Kafka-Consumer which listens for the events.
func (ka *KafkaFactory) consumer(name string, topics []string) (*consumer.Consumer, error) {
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

func (ka *KafkaFactory) EnsureConsumerIO(
	group string,
	topic string,
	enableErrors bool,
	cid uuuid.UUID,
) (<-chan model.KafkaResponse, error) {
	cidMapLock := &sync.RWMutex{}

	id := group + topic
	if cioStore[id] == nil {
		// Create Kafka Event-Consumer
		reqConsumer, err := ka.consumer(group, []string{topic})
		if err != nil {
			err = errors.Wrap(err, "Error Creating Response ConsumerGroup for Events")
			return nil, err
		}
		log.Println("Created Kafka Response ConsumerGroup")

		cioStore[id] = reqConsumer

		go func() {
			for msg := range reqConsumer.Messages() {
				reqConsumer.MarkOffset(msg, "")
				go handleKafkaConsumerMsg(msg, cidMap, cidMapLock)
			}
		}()
	}

	cidMapLock.RLock()
	topicCIDMap := cidMap[topic]
	cidMapLock.RUnlock()

	if topicCIDMap == nil {
		cidMapLock.Lock()
		cidMap[topic] = map[string]CIDSubAdapter{}
		cidMapLock.Unlock()

		cidMapLock.RLock()
		topicCIDMap = cidMap[topic]
		cidMapLock.RUnlock()
	}
	sa := newCIDSubAdapter(topicCIDMap, cid, cidMapLock)
	return sa.read(), nil
}

func handleKafkaConsumerMsg(
	msg *sarama.ConsumerMessage,
	cidMap map[string]map[string]CIDSubAdapter,
	cidMapLock *sync.RWMutex,
) {
	kr := model.KafkaResponse{}
	err := json.Unmarshal(msg.Value, &kr)
	if err != nil {
		err = errors.Wrap(err, "Error Unmarshalling Kafka-Message into Kafka-Response")
		log.Println(err)
		return
	}

	cidMapLock.RLock()
	topicCIDMap := cidMap[msg.Topic]
	cidMapLock.RUnlock()

	if topicCIDMap == nil {
		cidMapLock.Lock()
		cidMap[msg.Topic] = make(map[string]CIDSubAdapter)
		cidMapLock.Unlock()

		cidMapLock.RLock()
		topicCIDMap = cidMap[msg.Topic]
		cidMapLock.RUnlock()
	}
	krcid := kr.CorrelationID.String()

	topicMapLock := &sync.RWMutex{}

	topicMapLock.RLock()
	cidSubAdapter, exists := topicCIDMap[krcid]
	topicMapLock.RUnlock()

	if exists {
		topicMapLock.Lock()
		delete(topicCIDMap, krcid)
		topicMapLock.Unlock()

		cidSubAdapter.write(kr)
		return
	}

	cidSubAdapter = *newCIDSubAdapter(
		topicCIDMap,
		kr.CorrelationID,
		topicMapLock,
	)

	topicMapLock.Lock()
	topicCIDMap[krcid] = cidSubAdapter
	topicMapLock.Unlock()
}
