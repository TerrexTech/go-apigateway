package model

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type Report struct {
	ID           objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ReportID     uuuid.UUID        `bson:"report_id,omitempty" json:"report_id,omitempty"`
	RsCustomerID uuuid.UUID        `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	ItemID       uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
	DeviceID     uuuid.UUID        `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Timestamp    int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn       float64           `bson:"temp_in,omitempty" json:"temp_in,omitempty"`
	Humidity     float64           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene     float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDi     float64           `bson:"carbon_di,omitempty" json:"carbon_di,omitempty"`
	Version      int               `bson:"version,omitempty" json:"version,omitempty"`
}

type marshalReport struct {
	ID           objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ReportID     string            `bson:"report_id,omitempty" json:"report_id,omitempty"`
	RsCustomerID string            `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	ItemID       string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	DeviceID     string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Timestamp    int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TempIn       float64           `bson:"temp_in,omitempty" json:"temp_in,omitempty"`
	Humidity     float64           `bson:"humidity,omitempty" json:"humidity,omitempty"`
	Ethylene     float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDi     float64           `bson:"carbon_di,omitempty" json:"carbon_di,omitempty"`
	Version      int               `bson:"version,omitempty" json:"version,omitempty"`
}

func (r Report) MarshalBSON() ([]byte, error) {
	mr := &marshalReport{
		ID:        r.ID,
		Timestamp: r.Timestamp,
		Ethylene:  r.Ethylene,
		TempIn:    r.TempIn,
		Humidity:  r.Humidity,
		CarbonDi:  r.CarbonDi,
		Version:   r.Version,
	}

	if r.ReportID.String() != (uuuid.UUID{}).String() {
		mr.ReportID = r.ReportID.String()
	}

	if r.RsCustomerID.String() != (uuuid.UUID{}).String() {
		mr.RsCustomerID = r.RsCustomerID.String()
	}

	if r.ItemID.String() != (uuuid.UUID{}).String() {
		mr.ItemID = r.ItemID.String()
	}

	if r.DeviceID.String() != (uuuid.UUID{}).String() {
		mr.DeviceID = r.DeviceID.String()
	}
	return bson.Marshal(mr)
}

func (r *Report) MarshalJSON() ([]byte, error) {
	mr := &marshalReport{
		ID:        r.ID,
		Timestamp: r.Timestamp,
		Ethylene:  r.Ethylene,
		TempIn:    r.TempIn,
		Humidity:  r.Humidity,
		CarbonDi:  r.CarbonDi,
		Version:   r.Version,
	}

	if r.ReportID.String() != (uuuid.UUID{}).String() {
		mr.ReportID = r.ReportID.String()
	}

	if r.RsCustomerID.String() != (uuuid.UUID{}).String() {
		mr.RsCustomerID = r.RsCustomerID.String()
	}

	if r.ItemID.String() != (uuuid.UUID{}).String() {
		mr.ItemID = r.ItemID.String()
	}

	if r.DeviceID.String() != (uuuid.UUID{}).String() {
		mr.DeviceID = r.DeviceID.String()
	}

	return json.Marshal(mr)
}

func (r *Report) UnmarshalBSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if m["_id"] != nil {
		r.ID = m["_id"].(objectid.ObjectID)
	}

	if m["report_id"] != nil {
		r.ReportID, err = uuuid.FromString(m["report_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

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

	if m["rs_customer_id"] != nil {
		r.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
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

	return nil
}

func (r *Report) UnmarshalJSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if m["_id"] != nil {
		r.ID = m["_id"].(objectid.ObjectID)
	}

	if m["report_id"] != nil {
		r.ReportID, err = uuuid.FromString(m["report_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

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

	if m["rs_customer_id"] != nil {
		r.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
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

	return nil
}

// User represents a system-registered user.
// type Ethylene struct {
// 	ID       objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	UUID     uuuid.UUID        `bson:"uuid,omitempty" json:"uuid,omitempty"`
// 	ItemID   uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
// 	Ethylene float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
// 	SKU      int64             `bson:"sku,omitempty" json:"sku,omitempty"`
// 	Lot      string            `bson:"lot,omitempty" json:"lot,omitempty"`
// 	DeviceID uuuid.UUID        `bson:"device_id,omitempty" json:"device_id,omitempty"`
// }

// marshalUser is alternative format for User for convenient
// Marshalling/Unmarshalling operations.
// type marshalEthylene struct {
// 	ID       objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
// 	UUID     string            `bson:"uuid,omitempty" json:"uuid,omitempty"`
// 	ItemID   string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
// 	Ethylene float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
// 	SKU      int64             `bson:"sku,omitempty" json:"sku,omitempty"`
// 	Lot      string            `bson:"lot,omitempty" json:"lot,omitempty"`
// 	DeviceID string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
// }

// func (e Ethylene) MarshalBSON() ([]byte, error) {
// 	me := &marshalEthylene{
// 		ID:       e.ID,
// 		Ethylene: e.Ethylene,
// 		SKU:      e.SKU,
// 		Lot:      e.Lot,
// 	}

// 	if e.UUID.String() != (uuuid.UUID{}).String() {
// 		me.UUID = e.UUID.String()
// 	}

// 	if e.ItemID.String() != (uuuid.UUID{}).String() {
// 		me.ItemID = e.ItemID.String()
// 	}

// 	if e.DeviceID.String() != (uuuid.UUID{}).String() {
// 		me.DeviceID = e.DeviceID.String()
// 	}
// 	return bson.Marshal(me)
// }

// // MarshalJSON converts the current-user representation into its
// // JSON representation.
// func (e *Ethylene) MarshalJSON() ([]byte, error) {
// 	me := &marshalEthylene{
// 		ID:       e.ID,
// 		Ethylene: e.Ethylene,
// 		SKU:      e.SKU,
// 		Lot:      e.Lot,
// 	}

// 	if e.UUID.String() != (uuuid.UUID{}).String() {
// 		me.UUID = e.UUID.String()
// 	}

// 	if e.ItemID.String() != (uuuid.UUID{}).String() {
// 		me.ItemID = e.ItemID.String()
// 	}

// 	if e.DeviceID.String() != (uuuid.UUID{}).String() {
// 		me.DeviceID = e.DeviceID.String()
// 	}

// 	me.ID = e.ID
// 	// mu.UUID = u.UUID.String()

// 	return json.Marshal(me)
// }

// UnmarshalJSON converts the JSON representation of a User into the
// local User-struct.
func (e *Ethylene) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	log.Println(string(in))

	e.ID, err = objectid.FromHex(m["_id"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing ethylene _id")
		return err
	}

	e.UUID, err = uuuid.FromString(m["uuid"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing ethylene uuid")
		return err
	}

	e.ItemID, err = uuuid.FromString(m["item_id"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing ethylene item_id")
		return err
	}

	e.DeviceID, err = uuuid.FromString(m["device_id"].(string))
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error: Error parsing ethylene device_id")
		return err
	}

	if m["ethylene"] != nil {
		e.Ethylene = m["ethylene"].(float64)
	}

	if m["sku"] != nil {
		e.SKU = m["sku"].(int64)
	}

	if m["lot"] != nil {
		e.Lot = m["lot"].(string)
	}

	return nil
}
