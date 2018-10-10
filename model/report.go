package model

import (
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
	Version          int64             `bson:"version,omitempty" json:"version,omitempty"`
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
	Version          int64             `bson:"version,omitempty" json:"version,omitempty"`
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
		r.Version, ok = m["version"].(int64)
		if !ok {
			if versionType == reflect.Float64 {
				r.Version = int64(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				r.Version = int64(val)
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
		r.Version, ok = m["version"].(int64)
		if !ok {
			if versionType == reflect.Float64 {
				r.Version = int64(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				r.Version = int64(val)
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
