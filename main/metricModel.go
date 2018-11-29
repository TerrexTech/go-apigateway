package main

import (
	"encoding/json"

	util "github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// AggregateID is the global AggregateID for Inventory Aggregate.

//Metric struct
type Metric struct {
	ID            objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MetricID      string            `bson:"metricID,omitempty" json:"metricID,omitempty"`
	ItemID        string            `bson:"itemID,omitempty" json:"itemID,omitempty"`
	DeviceID      string            `bson:"deviceID,omitempty" json:"deviceID,omitempty"`
	Timestamp     int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn        float64           `bson:"tempIn,omitempty" json:"tempIn,omitempty"`
	Humidity      float64           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene      float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float64           `bson:"carbonDioxide,omitempty" json:"carbonDioxide,omitempty"`
	SKU           string            `bson:"sku,omitempty" json:"sku,omitempty"`
	Name          string            `bson:"name,omitempty" json:"name,omitempty"`
	Lot           string            `bson:"lot,omitempty" json:"lot,omitempty"`
}

// MarshalBSON returns bytes of BSON-type.
func (m *Metric) MarshalBSON() ([]byte, error) {
	mm := map[string]interface{}{
		"metricID":      m.MetricID,
		"itemID":        m.ItemID,
		"deviceID":      m.DeviceID,
		"timestamp":     m.Timestamp,
		"tempIn":        m.TempIn,
		"humidity":      m.Humidity,
		"ethylene":      m.Ethylene,
		"carbonDioxide": m.CarbonDioxide,
		"sku":           m.SKU,
		"name":          m.Name,
		"lot":           m.Lot,
	}

	if m.ID != objectid.NilObjectID {
		mm["_id"] = m.ID
	}

	return bson.Marshal(mm)
}

// MarshalJSON returns bytes of JSON-type.
func (m *Metric) MarshalJSON() ([]byte, error) {
	mm := map[string]interface{}{
		"metricID":      m.MetricID,
		"itemID":        m.ItemID,
		"deviceID":      m.DeviceID,
		"timestamp":     m.Timestamp,
		"tempIn":        m.TempIn,
		"humidity":      m.Humidity,
		"ethylene":      m.Ethylene,
		"carbonDioxide": m.CarbonDioxide,
		"sku":           m.SKU,
		"name":          m.Name,
		"lot":           m.Lot,
	}

	if m.ID != objectid.NilObjectID {
		mm["_id"] = m.ID.Hex()
	}

	return json.Marshal(mm)
}

// UnmarshalBSON returns BSON-type from bytes.
func (m *Metric) UnmarshalBSON(in []byte) error {
	metMap := make(map[string]interface{})
	err := bson.Unmarshal(in, metMap)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalBSON Error")
		return err
	}

	err = m.unmarshalFromMap(metMap)
	return err
}

// UnmarshalJSON returns JSON-type from bytes.
func (m *Metric) UnmarshalJSON(in []byte) error {
	metMap := make(map[string]interface{})
	err := json.Unmarshal(in, &metMap)
	if err != nil {
		err = errors.Wrap(err, "UnmarshalBSON Error")
		return err
	}

	err = m.unmarshalFromMap(metMap)
	return err
}

func (m *Metric) unmarshalFromMap(metMap map[string]interface{}) error {
	var err error
	var assertOK bool

	if metMap["_id"] != nil {
		m.ID, assertOK = metMap["_id"].(objectid.ObjectID)
		if !assertOK {
			m.ID, err = objectid.FromHex(metMap["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if metMap["metricID"] != nil {
		m.MetricID, assertOK = metMap["metricID"].(string)
		if !assertOK {
			return errors.New("error asserting MetricID")
		}
		// m.MetricID, err = uuuid.FromString(metricIDstr)
		// if err != nil {
		// 	err = errors.Wrap(err, "Error while asserting metricID")
		// 	return err
		// }
	}

	if metMap["itemID"] != nil {
		m.ItemID, assertOK = metMap["itemID"].(string)
		if assertOK {
			return errors.New("error asserting ItemID")
		}
		// m.ItemID, err = uuuid.FromString(itemIDStr)
		// if err != nil {
		// 	err = errors.Wrap(err, "Error while asserting itemID")
		// 	return err
		// }
	}

	if metMap["deviceID"] != nil {
		m.DeviceID, assertOK = metMap["deviceID"].(string)
		if !assertOK {
			return errors.New("error asserting DeviceID")
		}
		// m.DeviceID, err = uuuid.FromString(deviceIDStr)
		// if err != nil {
		// 	err = errors.Wrap(err, "Error while asserting deviceID")
		// 	return err
		// }
	}

	if metMap["sku"] != nil {
		m.SKU, assertOK = metMap["sku"].(string)
		if !assertOK {
			return errors.New("error asserting SKU")
		}
	}

	if metMap["lot"] != nil {
		m.Lot, assertOK = metMap["lot"].(string)
		if !assertOK {
			return errors.New("error asserting Lot")
		}
	}

	if metMap["name"] != nil {
		m.Name, assertOK = metMap["name"].(string)
		if !assertOK {
			return errors.New("error asserting Name")
		}
	}

	if metMap["timestamp"] != nil {
		m.Timestamp, err = util.AssertInt64(metMap["timestamp"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Timestamp")
			return err
		}
	}

	if metMap["tempIn"] != nil {
		m.TempIn, err = util.AssertFloat64(metMap["tempIn"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting tempIn")
			return err
		}
	}

	if metMap["humidity"] != nil {
		m.Humidity, err = util.AssertFloat64(metMap["humidity"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting humidity")
			return err
		}
	}

	if metMap["ethylene"] != nil {
		m.Ethylene, err = util.AssertFloat64(metMap["ethylene"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting ethylene")
			return err
		}
	}

	if metMap["carbonDioxide"] != nil {
		m.CarbonDioxide, err = util.AssertFloat64(metMap["carbonDioxide"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting carbonDioxide")
			return err
		}
	}

	return nil
}
