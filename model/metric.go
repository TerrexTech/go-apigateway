package model

import (
	"encoding/json"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type Metric struct {
	ID            objectid.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	ItemID        uuuid.UUID        `bson:"itemID,omitempty" json:"itemID,omitempty"`
	DeviceID      uuuid.UUID        `bson:"deviceID,omitempty" json:"deviceID,omitempty"`
	Timestamp     int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn        float32           `bson:"tempIn,omitempty" json:"tempIn,omitempty"`
	Humidity      float32           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene      float32           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float32           `bson:"carbonDioxide,omitempty" json:"carbonDioxide,omitempty"`
	Version       int64             `bson:"version,omitempty" json:"version,omitempty"`
}

type marshalMetric struct {
	ID            objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RsCustomerID  string            `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	ItemID        string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	DeviceID      string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Timestamp     int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn        float32           `bson:"temp_in,omitempty" json:"temp_in,omitempty"`
	Humidity      float32           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene      float32           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float32           `bson:"carbon_dioxide,omitempty" json:"carbon_dioxide,omitempty"`
	Version       int64             `bson:"version,omitempty" json:"version,omitempty"`
}

func (m *Metric) MarshalJSON() ([]byte, error) {
	mm := &marshalMetric{
		ID:            m.ID,
		Timestamp:     m.Timestamp,
		Ethylene:      m.Ethylene,
		TempIn:        m.TempIn,
		Humidity:      m.Humidity,
		CarbonDioxide: m.CarbonDioxide,
		Version:       m.Version,
	}

	if m.ItemID.String() != (uuuid.UUID{}).String() {
		mm.ItemID = m.ItemID.String()
	}

	if m.DeviceID.String() != (uuuid.UUID{}).String() {
		mm.DeviceID = m.DeviceID.String()
	}

	return json.Marshal(mm)
}

func (m *Metric) UnmarshalJSON(in []byte) error {
	metricMap := make(map[string]interface{})
	err := json.Unmarshal(in, &metricMap)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if metricMap["_id"] != nil {
		m.ID = metricMap["_id"].(objectid.ObjectID)
	}

	if metricMap["item_id"] != nil {
		m.ItemID, err = uuuid.FromString(metricMap["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if metricMap["device_id"] != nil {
		m.DeviceID, err = uuuid.FromString(metricMap["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	m.Timestamp = metricMap["timestamp"].(int64)
	m.TempIn = metricMap["temp_in"].(float32)
	m.Humidity = metricMap["humidity"].(float32)
	m.Ethylene = metricMap["ethylene"].(float32)
	m.CarbonDioxide = metricMap["carbon_dioxide"].(float32)
	m.Version = metricMap["version"].(int64)

	return nil
}
