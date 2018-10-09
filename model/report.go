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

type Report struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID           uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
	ReportID         uuuid.UUID        `bson:"report_id,omitempty" json:"report_id,omitempty"`
	RsCustomerID     uuuid.UUID        `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	ReportType       string            `bson:"report_type,omitempty" json:"report_type,omitempty"`
	Version          int               `bson:"version,omitempty" json:"version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

type marshalReport struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID           string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	ReportID         string            `bson:"report_id,omitempty" json:"report_id,omitempty"`
	RsCustomerID     string            `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	ReportType       string            `bson:"report_type,omitempty" json:"report_type,omitempty"`
	Version          int               `bson:"version,omitempty" json:"version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

func (r Report) MarshalBSON() ([]byte, error) {
	mr := &marshalReport{
		ID:               r.ID,
		ReportType:       r.ReportType,
		Timestamp:        r.Timestamp,
		Version:          r.Version,
		AggregateID:      r.AggregateID,
		AggregateVersion: r.AggregateVersion,
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

	return bson.Marshal(mr)
}

func (r Report) MarshalJSON() ([]byte, error) {
	mr := &marshalReport{
		ID:               r.ID,
		ReportType:       r.ReportType,
		Timestamp:        r.Timestamp,
		Version:          r.Version,
		AggregateID:      r.AggregateID,
		AggregateVersion: r.AggregateVersion,
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

	return bson.Marshal(mr)
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

	if m["rs_customer_id"] != nil {
		r.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
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

	if m["report_type"] != nil {
		r.ReportType = m["report_type"].(string)
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

	if m["rs_customer_id"] != nil {
		r.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
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

	if m["report_type"] != nil {
		r.ReportType = m["report_type"].(string)
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

	if m["_id"] != nil {
		r.ID = m["_id"].(objectid.ObjectID)
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

	if m["_id"] != nil {
		r.ID = m["_id"].(objectid.ObjectID)
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

type Inventory struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID           uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
	UPC              int64             `bson:"upc,omitempty" json:"upc,omitempty"`
	SKU              int64             `bson:"sku,omitempty" json:"sku,omitempty"`
	Name             string            `bson:"name,omitempty" json:"name,omitempty"`
	Origin           string            `bson:"origin,omitempty" json:"origin,omitempty"`
	DeviceID         uuuid.UUID        `bson:"device_id,omitempty" json:"device_id,omitempty"`
	TotalWeight      float64           `bson:"total_weight,omitempty" json:"total_weight,omitempty"`
	Price            float64           `bson:"price,omitempty" json:"price,omitempty"`
	Location         string            `bson:"location,omitempty" json:"location,omitempty"`
	DateArrived      int64             `bson:"date_arrived,omitempty" json:"date_arrived,omitempty"`
	ExpiryDate       int64             `bson:"expiry_date,omitempty" json:"expiry_date,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	RsCustomerID     uuuid.UUID        `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	WasteWeight      float64           `bson:"waste_weight,omitempty" json:"waste_weight,omitempty"`
	DonateWeight     float64           `bson:"donate_weight,omitempty" json:"donate_weight,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	DateSold         int64             `bson:"date_sold,omitempty" json:"date_sold,omitempty"`
	SalePrice        float64           `bson:"sale_price,omitempty" json:"sale_price,omitempty"`
	SoldWeight       float64           `bson:"sold_weight,omitempty" json:"sold_weight,omitempty"`
	ProdQuantity     int64             `bson:"prod_quantity,omitempty" json:"prod_quantity,omitempty"`
}

type marshalInventory struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID           string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	UPC              int64             `bson:"upc,omitempty" json:"upc,omitempty"`
	SKU              int64             `bson:"sku,omitempty" json:"sku,omitempty"`
	Name             string            `bson:"name,omitempty" json:"name,omitempty"`
	Origin           string            `bson:"origin,omitempty" json:"origin,omitempty"`
	DeviceID         string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
	TotalWeight      float64           `bson:"total_weight,omitempty" json:"total_weight,omitempty"`
	Price            float64           `bson:"price,omitempty" json:"price,omitempty"`
	Location         string            `bson:"location,omitempty" json:"location,omitempty"`
	DateArrived      int64             `bson:"date_arrived,omitempty" json:"date_arrived,omitempty"`
	ExpiryDate       int64             `bson:"expiry_date,omitempty" json:"expiry_date,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	RsCustomerID     string            `bson:"rs_customer_id,omitempty" json:"rs_customer_id,omitempty"`
	WasteWeight      float64           `bson:"waste_weight,omitempty" json:"waste_weight,omitempty"`
	DonateWeight     float64           `bson:"donate_weight,omitempty" json:"donate_weight,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
	AggregateID      int8              `bson:"aggregate_id,omitempty" json:"aggregate_id,omitempty"`
	DateSold         int64             `bson:"date_sold,omitempty" json:"date_sold,omitempty"`
	SalePrice        float64           `bson:"sale_price,omitempty" json:"sale_price,omitempty"`
	SoldWeight       float64           `bson:"sold_weight,omitempty" json:"sold_weight,omitempty"`
	ProdQuantity     int64             `bson:"prod_quantity,omitempty" json:"prod_quantity,omitempty"`
}

func (i *Inventory) MarshalJSON() ([]byte, error) {
	in := &marshalInventory{
		UPC:              i.UPC,
		SKU:              i.SKU,
		Name:             i.Name,
		Origin:           i.Origin,
		TotalWeight:      i.TotalWeight,
		Price:            i.Price,
		Location:         i.Location,
		DateArrived:      i.DateArrived,
		ExpiryDate:       i.ExpiryDate,
		Timestamp:        i.Timestamp,
		WasteWeight:      i.WasteWeight,
		DonateWeight:     i.DonateWeight,
		AggregateVersion: i.AggregateVersion,
		AggregateID:      i.AggregateID,
		DateSold:         i.DateSold,
		SalePrice:        i.SalePrice,
		SoldWeight:       i.SoldWeight,
	}

	if i.ItemID.String() != (uuuid.UUID{}).String() {
		in.ItemID = i.ItemID.String()
	}
	if i.DeviceID.String() != (uuuid.UUID{}).String() {
		in.DeviceID = i.DeviceID.String()
	}
	if i.RsCustomerID.String() != (uuuid.UUID{}).String() {
		in.RsCustomerID = i.RsCustomerID.String()
	}

	return json.Marshal(in)
}

func (i Inventory) MarshalBSON() ([]byte, error) {
	in := &marshalInventory{
		UPC:              i.UPC,
		SKU:              i.SKU,
		Name:             i.Name,
		Origin:           i.Origin,
		TotalWeight:      i.TotalWeight,
		Price:            i.Price,
		Location:         i.Location,
		DateArrived:      i.DateArrived,
		ExpiryDate:       i.ExpiryDate,
		Timestamp:        i.Timestamp,
		WasteWeight:      i.WasteWeight,
		DonateWeight:     i.DonateWeight,
		AggregateVersion: i.AggregateVersion,
		AggregateID:      i.AggregateID,
		DateSold:         i.DateSold,
		SalePrice:        i.SalePrice,
		SoldWeight:       i.SoldWeight,
	}

	if i.ItemID.String() != (uuuid.UUID{}).String() {
		in.ItemID = i.ItemID.String()
	}

	if i.DeviceID.String() != (uuuid.UUID{}).String() {
		in.DeviceID = i.DeviceID.String()
	}
	if i.RsCustomerID.String() != (uuuid.UUID{}).String() {
		in.RsCustomerID = i.RsCustomerID.String()
	}

	return bson.Marshal(in)
}

func (i *Inventory) UnmarshalBSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if m["_id"] != nil {
		i.ID = m["_id"].(objectid.ObjectID)
	}

	// log.Println((m["item_id"].(map[string]interface{}))["uuid"])

	if m["item_id"] != nil {
		i.ItemID, err = uuuid.FromString(m["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["upc"] != nil {
		upcType := reflect.TypeOf(m["upc"]).Kind()
		i.UPC, ok = m["upc"].(int64)
		if !ok {
			if upcType == reflect.Float64 {
				i.UPC = int64(m["upc"].(float64))
			} else {
				val, _ := strconv.Atoi((m["upc"]).(string))
				i.UPC = int64(val)
			}

		}
	}

	if m["sku"] != nil {
		skuType := reflect.TypeOf(m["sku"]).Kind()
		i.SKU, ok = m["sku"].(int64)
		if !ok {
			if skuType == reflect.Float64 {
				i.SKU = int64(m["sku"].(float64))
			} else {
				val, _ := strconv.Atoi((m["sku"]).(string))
				i.SKU = int64(val)
			}

		}
	}

	if m["device_id"] != nil {
		i.DeviceID, err = uuuid.FromString(m["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["rs_customer_id"] != nil {
		i.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
	}
	if err != nil {
		err = errors.Wrap(err, "Error parsing DeviceID for inventory")
		return err
	}

	if m["name"] != nil {
		i.Name = m["name"].(string)
	}

	if m["origin"] != nil {
		i.Origin = m["origin"].(string)
	}

	if m["total_weight"] != nil {
		totalWeightType := reflect.TypeOf(m["total_weight"]).Kind()
		i.TotalWeight, ok = m["total_weight"].(float64)
		if !ok {
			if totalWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["total_weight"]).(string))
				i.TotalWeight = float64(val)
			}
		}
	}

	if m["price"] != nil {
		priceType := reflect.TypeOf(m["price"]).Kind()
		i.Price, ok = m["price"].(float64)
		if !ok {
			if priceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["price"]).(string))
				i.Price = float64(val)
			}
		}
	}

	if m["location"] != nil {
		i.Location = m["location"].(string)
	}

	if m["date_arrived"] != nil {
		datearrivedType := reflect.TypeOf(m["date_arrived"]).Kind()
		i.DateArrived, ok = m["date_arrived"].(int64)
		if !ok {
			if datearrivedType == reflect.Float64 {
				i.DateArrived = int64(m["date_arrived"].(float64))
			} else {
				val, _ := strconv.Atoi((m["date_arrived"]).(string))
				i.DateArrived = int64(val)
			}

		}
	}

	if m["expiry_date"] != nil {
		expiryDateType := reflect.TypeOf(m["expiry_date"]).Kind()
		i.ExpiryDate, ok = m["expiry_date"].(int64)
		if !ok {
			if expiryDateType == reflect.Float64 {
				i.ExpiryDate = int64(m["expiry_date"].(float64))
			} else {
				val, _ := strconv.Atoi((m["expiry_date"]).(string))
				i.ExpiryDate = int64(val)
			}

		}
	}

	if m["timestamp"] != nil {
		timestampType := reflect.TypeOf(m["timestamp"]).Kind()
		i.Timestamp, ok = m["timestamp"].(int64)
		if !ok {
			if timestampType == reflect.Float64 {
				i.Timestamp = int64(m["timestamp"].(float64))
			} else {
				val, _ := strconv.Atoi((m["timestamp"]).(string))
				i.Timestamp = int64(val)
			}
		}
	}

	if m["date_sold"] != nil {
		datesoldType := reflect.TypeOf(m["date_sold"]).Kind()
		i.DateSold, ok = m["date_sold"].(int64)
		if !ok {
			if datesoldType == reflect.Float64 {
				i.DateSold = int64(m["date_sold"].(float64))
			} else {
				val, _ := strconv.Atoi((m["date_sold"]).(string))
				i.DateSold = int64(val)
			}
		}
	}

	if m["waste_weight"] != nil {
		wasteWeightType := reflect.TypeOf(m["waste_weight"]).Kind()
		i.WasteWeight, ok = m["waste_weight"].(float64)
		if !ok {
			if wasteWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["waste_weight"]).(string))
				i.WasteWeight = float64(val)
			}
		}
	}

	if m["donate_weight"] != nil {
		donateWeightType := reflect.TypeOf(m["donate_weight"]).Kind()
		i.DonateWeight, ok = m["donate_weight"].(float64)
		if !ok {
			if donateWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["donate_weight"]).(string))
				i.DonateWeight = float64(val)
			}
		}
	}

	if m["aggregate_version"] != nil {
		aggregateVersionType := reflect.TypeOf(m["aggregate_version"]).Kind()
		i.AggregateVersion, ok = m["aggregate_version"].(int64)
		if !ok {
			if aggregateVersionType != reflect.Int64 {
				val, _ := strconv.Atoi((m["aggregate_version"]).(string))
				i.AggregateVersion = int64(val)
			}
		}
	}

	if m["aggregate_id"] != nil {
		// i.AggregateID = m["aggregate_id"].(int8)
	}

	if m["sale_price"] != nil {
		salePriceType := reflect.TypeOf(m["sale_price"]).Kind()
		i.SalePrice, ok = m["sale_price"].(float64)
		if !ok {
			if salePriceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["sale_price"]).(string))
				i.SalePrice = float64(val)
			}
		}
	}

	if m["sold_weight"] != nil {
		soldWeightType := reflect.TypeOf(m["sold_weight"]).Kind()
		i.SoldWeight = m["sold_weight"].(float64)
		if !ok {
			if soldWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["sold_weight"]).(string))
				i.SoldWeight = float64(val)
			}
		}
	}

	return nil
}

func (i *Inventory) UnmarshalJSON(in []byte) error {
	var ok bool
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if m["_id"] != nil {
		i.ID = m["_id"].(objectid.ObjectID)
	}

	if m["item_id"] != nil {
		i.ItemID, err = uuuid.FromString(m["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["upc"] != nil {
		upcType := reflect.TypeOf(m["upc"]).Kind()
		i.UPC, ok = m["upc"].(int64)
		if !ok {
			if upcType == reflect.Float64 {
				i.UPC = int64(m["upc"].(float64))
			} else {
				val, _ := strconv.Atoi((m["upc"]).(string))
				i.UPC = int64(val)
			}
		}
	}

	if m["sku"] != nil {
		skuType := reflect.TypeOf(m["sku"]).Kind()
		i.SKU, ok = m["sku"].(int64)
		if !ok {
			if skuType == reflect.Float64 {
				i.SKU = int64(m["sku"].(float64))
			} else {
				val, _ := strconv.Atoi((m["sku"]).(string))
				i.SKU = int64(val)
			}

		}
	}

	if m["device_id"] != nil {
		i.DeviceID, err = uuuid.FromString(m["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["rs_customer_id"] != nil {
		i.RsCustomerID, err = uuuid.FromString(m["rs_customer_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["name"] != nil {
		i.Name = m["name"].(string)
	}

	if m["origin"] != nil {
		i.Origin = m["origin"].(string)
	}

	if m["total_weight"] != nil {
		totalWeightType := reflect.TypeOf(m["total_weight"]).Kind()
		i.TotalWeight, ok = m["total_weight"].(float64)
		if !ok {
			if totalWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["total_weight"]).(string))
				i.TotalWeight = float64(val)
			}
		}
	}

	if m["price"] != nil {
		priceType := reflect.TypeOf(m["price"]).Kind()
		i.Price, ok = m["price"].(float64)
		if !ok {
			if priceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["price"]).(string))
				i.Price = float64(val)
			}
		}
	}

	if m["location"] != nil {
		i.Location = m["location"].(string)
	}

	if m["date_arrived"] != nil {
		datearrivedType := reflect.TypeOf(m["date_arrived"]).Kind()
		i.DateArrived, ok = m["date_arrived"].(int64)
		if !ok {
			if datearrivedType == reflect.Float64 {
				i.DateArrived = int64(m["date_arrived"].(float64))
			} else {
				val, _ := strconv.Atoi((m["date_arrived"]).(string))
				i.DateArrived = int64(val)
			}

		}
	}

	if m["expiry_date"] != nil {
		expiryDateType := reflect.TypeOf(m["expiry_date"]).Kind()
		i.ExpiryDate, ok = m["expiry_date"].(int64)
		if !ok {
			if expiryDateType == reflect.Float64 {
				i.ExpiryDate = int64(m["expiry_date"].(float64))
			} else {
				val, _ := strconv.Atoi((m["expiry_date"]).(string))
				i.ExpiryDate = int64(val)
			}

		}
	}

	if m["timestamp"] != nil {
		timestampType := reflect.TypeOf(m["timestamp"]).Kind()
		i.Timestamp, ok = m["timestamp"].(int64)
		if !ok {
			if timestampType == reflect.Float64 {
				i.Timestamp = int64(m["timestamp"].(float64))
			} else {
				val, _ := strconv.Atoi((m["timestamp"]).(string))
				i.Timestamp = int64(val)
			}
		}
	}

	if m["date_sold"] != nil {
		datesoldType := reflect.TypeOf(m["date_sold"]).Kind()
		i.DateSold, ok = m["date_sold"].(int64)
		if !ok {
			if datesoldType == reflect.Float64 {
				i.DateSold = int64(m["date_sold"].(float64))
			} else {
				val, _ := strconv.Atoi((m["date_sold"]).(string))
				i.DateSold = int64(val)
			}
		}
	}

	if m["waste_weight"] != nil {
		wasteWeightType := reflect.TypeOf(m["waste_weight"]).Kind()
		i.WasteWeight, ok = m["waste_weight"].(float64)
		if !ok {
			if wasteWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["waste_weight"]).(string))
				i.WasteWeight = float64(val)
			}
		}
	}

	if m["donate_weight"] != nil {
		donateWeightType := reflect.TypeOf(m["donate_weight"]).Kind()
		i.DonateWeight, ok = m["donate_weight"].(float64)
		if !ok {
			if donateWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["donate_weight"]).(string))
				i.DonateWeight = float64(val)
			}
		}
	}

	if m["aggregate_version"] != nil {
		aggregateVersionType := reflect.TypeOf(m["aggregate_version"]).Kind()
		i.AggregateVersion, ok = m["aggregate_version"].(int64)
		if !ok {
			if aggregateVersionType != reflect.Int64 {
				val, _ := strconv.Atoi((m["aggregate_version"]).(string))
				i.AggregateVersion = int64(val)
			}
		}
	}

	if m["aggregate_id"] != nil {
		// i.AggregateID = m["aggregate_id"].(int8)
	}

	if m["sale_price"] != nil {
		salePriceType := reflect.TypeOf(m["sale_price"]).Kind()
		i.SalePrice, ok = m["sale_price"].(float64)
		if !ok {
			if salePriceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["sale_price"]).(string))
				i.SalePrice = float64(val)
			}
		}
	}

	if m["sold_weight"] != nil {
		soldWeightType := reflect.TypeOf(m["sold_weight"]).Kind()
		i.SoldWeight = m["sold_weight"].(float64)
		if !ok {
			if soldWeightType != reflect.Float64 {
				val, _ := strconv.Atoi((m["sold_weight"]).(string))
				i.SoldWeight = float64(val)
			}
		}
	}
	return nil
}
