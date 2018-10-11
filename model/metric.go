package model

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type Metric struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID           uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
	DeviceID         uuuid.UUID        `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn           float64           `bson:"temp_in,omitempty" json:"temp_in,omitempty"`
	Humidity         float64           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene         float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDi         float64           `bson:"carbon_di,omitempty" json:"carbon_di,omitempty"`
	Version          int               `bson:"version,omitempty" json:"version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

type marshalMetric struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	RsCustomerID     string            `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	ItemID           string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	DeviceID         string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn           float64           `bson:"temp_in,omitempty" json:"temp_in,omitempty"`
	Humidity         float64           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene         float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDi         float64           `bson:"carbon_di,omitempty" json:"carbon_di,omitempty"`
	Version          int               `bson:"version,omitempty" json:"version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

func (m Metric) MarshalBSON() ([]byte, error) {
	mm := &marshalMetric{
		ID:               m.ID,
		Timestamp:        m.Timestamp,
		Ethylene:         m.Ethylene,
		TempIn:           m.TempIn,
		Humidity:         m.Humidity,
		CarbonDi:         m.CarbonDi,
		Version:          m.Version,
		AggregateID:      m.AggregateID,
		AggregateVersion: m.AggregateVersion,
	}

	if m.ItemID.String() != (uuuid.UUID{}).String() {
		mm.ItemID = m.ItemID.String()
	}

	if m.DeviceID.String() != (uuuid.UUID{}).String() {
		mm.DeviceID = m.DeviceID.String()
	}
	return bson.Marshal(mm)
}

func (m *Metric) MarshalJSON() ([]byte, error) {
	mm := &marshalMetric{
		ID:               m.ID,
		Timestamp:        m.Timestamp,
		Ethylene:         m.Ethylene,
		TempIn:           m.TempIn,
		Humidity:         m.Humidity,
		CarbonDi:         m.CarbonDi,
		Version:          m.Version,
		AggregateID:      m.AggregateID,
		AggregateVersion: m.AggregateVersion,
	}

	if m.ItemID.String() != (uuuid.UUID{}).String() {
		mm.ItemID = m.ItemID.String()
	}

	if m.DeviceID.String() != (uuuid.UUID{}).String() {
		mm.DeviceID = m.DeviceID.String()
	}

	return json.Marshal(mm)
}

func (r *Metric) UnmarshalBSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	// if m["_id"] != nil {
	// 	r.ID = m["_id"].(objectid.ObjectID)
	// }

	if m["item_id"] != nil {
		r.ItemID, err = uuuid.FromString(m["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["device_id"] != nil {
		r.DeviceID, err = uuuid.FromString(m["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["temp_in"] != nil {
		tempInType := reflect.TypeOf(m["temp_in"]).Kind()
		r.TempIn, ok = m["temp_in"].(float64)
		if !ok {
			if tempInType != reflect.Float64 {
				val, _ := strconv.Atoi((m["temp_in"]).(string))
				r.TempIn = float64(val)
			}
		}
	}

	if m["humidity"] != nil {
		humidityType := reflect.TypeOf(m["humidity"]).Kind()
		r.Humidity, ok = m["humidity"].(float64)
		if !ok {
			if humidityType != reflect.Float64 {
				val, _ := strconv.Atoi((m["humidity"]).(string))
				r.Humidity = float64(val)
			}
		}
	}

	if m["ethylene"] != nil {
		ethyleneType := reflect.TypeOf(m["ethylene"]).Kind()
		r.Ethylene, ok = m["ethylene"].(float64)
		if !ok {
			if ethyleneType != reflect.Float64 {
				val, _ := strconv.Atoi((m["ethylene"]).(string))
				r.Ethylene = float64(val)
			}
		}
	}

	if m["carbon_di"] != nil {
		carbonType := reflect.TypeOf(m["carbon_di"]).Kind()
		r.CarbonDi, ok = m["carbon_di"].(float64)
		if !ok {
			if carbonType != reflect.Float64 {
				val, _ := strconv.Atoi((m["carbon_di"]).(string))
				r.CarbonDi = float64(val)
			}
		}
	}

	if m["timestamp"] != nil {
		timestampType := reflect.TypeOf(m["timestamp"]).Kind()
		r.Timestamp, ok = m["timestamp"].(int64)
		if !ok {
			if timestampType == reflect.Float64 {
				r.Timestamp = int64(m["timestamp"].(float64))
			} else {
				val, _ := strconv.Atoi((m["timestamp"]).(string))
				r.Timestamp = int64(val)
			}
		}
	}

	if m["version"] != nil {
		versionType := reflect.TypeOf(m["version"]).Kind()
		r.Version, ok = m["version"].(int)
		if !ok {
			if versionType == reflect.Float64 {
				r.Version = int(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				r.Version = int(val)
			}
		}
	}

	if m["aggregate_id"] != nil {
		aggregateIdType := reflect.TypeOf(m["aggregate_id"]).Kind()
		r.AggregateID, ok = m["aggregate_id"].(int8)
		if !ok {
			if aggregateIdType != reflect.Int8 {
				val, _ := strconv.Atoi((m["aggregate_id"]).(string))
				r.AggregateID = int8(val)
			}
		}
	}

	if m["aggregate_version"] != nil {
		aggregateVersionType := reflect.TypeOf(m["aggregate_version"]).Kind()
		r.AggregateVersion, ok = m["aggregate_version"].(int64)
		if !ok {
			if aggregateVersionType != reflect.Int64 {
				val, _ := strconv.Atoi((m["aggregate_version"]).(string))
				r.AggregateVersion = int64(val)
			}
		}
	}

	return nil
}

func (r *Metric) UnmarshalJSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	// if m["_id"] != nil {
	// 	r.ID = m["_id"].(objectid.ObjectID)
	// }

	if m["item_id"] != nil {
		r.ItemID, err = uuuid.FromString(m["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["device_id"] != nil {
		r.DeviceID, err = uuuid.FromString(m["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["temp_in"] != nil {
		tempInType := reflect.TypeOf(m["temp_in"]).Kind()
		r.TempIn, ok = m["temp_in"].(float64)
		if !ok {
			if tempInType != reflect.Float64 {
				val, _ := strconv.Atoi((m["temp_in"]).(string))
				r.TempIn = float64(val)
			}
		}
	}

	if m["humidity"] != nil {
		humidityType := reflect.TypeOf(m["humidity"]).Kind()
		r.Humidity, ok = m["humidity"].(float64)
		if !ok {
			if humidityType != reflect.Float64 {
				val, _ := strconv.Atoi((m["humidity"]).(string))
				r.Humidity = float64(val)
			}
		}
	}

	if m["ethylene"] != nil {
		ethyleneType := reflect.TypeOf(m["ethylene"]).Kind()
		r.Ethylene, ok = m["ethylene"].(float64)
		if !ok {
			if ethyleneType != reflect.Float64 {
				val, _ := strconv.Atoi((m["ethylene"]).(string))
				r.Ethylene = float64(val)
			}
		}
	}

	if m["carbon_di"] != nil {
		carbonType := reflect.TypeOf(m["carbon_di"]).Kind()
		r.CarbonDi, ok = m["carbon_di"].(float64)
		if !ok {
			if carbonType != reflect.Float64 {
				val, _ := strconv.Atoi((m["carbon_di"]).(string))
				r.CarbonDi = float64(val)
			}
		}
	}

	if m["timestamp"] != nil {
		timestampType := reflect.TypeOf(m["timestamp"]).Kind()
		r.Timestamp, ok = m["timestamp"].(int64)
		if !ok {
			if timestampType == reflect.Float64 {
				r.Timestamp = int64(m["timestamp"].(float64))
			} else {
				val, _ := strconv.Atoi((m["timestamp"]).(string))
				r.Timestamp = int64(val)
			}
		}
	}

	if m["version"] != nil {
		versionType := reflect.TypeOf(m["version"]).Kind()
		r.Version, ok = m["version"].(int)
		if !ok {
			if versionType == reflect.Float64 {
				r.Version = int(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				r.Version = int(val)
			}
		}
	}

	if m["aggregate_id"] != nil {
		aggregateIdType := reflect.TypeOf(m["aggregate_id"]).Kind()
		r.AggregateID, ok = m["aggregate_id"].(int8)
		if !ok {
			if aggregateIdType != reflect.Int8 {
				val, _ := strconv.Atoi((m["aggregate_id"]).(string))
				r.AggregateID = int8(val)
			}
		}
	}

	if m["aggregate_version"] != nil {
		aggregateVersionType := reflect.TypeOf(m["aggregate_version"]).Kind()
		r.AggregateVersion, ok = m["aggregate_version"].(int64)
		if !ok {
			if aggregateVersionType != reflect.Int64 {
				val, _ := strconv.Atoi((m["aggregate_version"]).(string))
				r.AggregateVersion = int64(val)
			}
		}
	}

	return nil
}
