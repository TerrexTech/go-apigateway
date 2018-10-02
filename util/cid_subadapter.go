package util

import (
	"sync"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
)

// CIDStore is the global CorrelationID (CID) storage, intended to be used for operations
// that require sharing CIDs across instances (go routines)
type CIDSubAdapter struct {
	cid      uuuid.UUID
	cidMap   map[string]CIDSubAdapter
	readChan chan model.KafkaResponse
}

func newCIDSubAdapter(
	topicCIDMap map[string]CIDSubAdapter,
	cid uuuid.UUID,
	topicMapLock *sync.RWMutex,
) *CIDSubAdapter {
	readChan := make(chan model.KafkaResponse, 1)

	ca := CIDSubAdapter{
		cid:      cid,
		cidMap:   topicCIDMap,
		readChan: readChan,
	}
	topicMapLock.Lock()
	topicCIDMap[cid.String()] = ca
	topicMapLock.Unlock()
	return &ca
}

func (c *CIDSubAdapter) read() <-chan model.KafkaResponse {
	return c.readChan
}

func (c *CIDSubAdapter) write(kr model.KafkaResponse) {
	c.readChan <- kr
}
