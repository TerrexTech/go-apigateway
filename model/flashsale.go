package model

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/TerrexTech/uuuid"
	bson "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type Flash struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FlashID          uuuid.UUID        `bson:"flash_id,omitempty" json:"flash_id,omitempty"`
	ItemID           uuuid.UUID        `bson:"item_id,omitempty" json:"item_id,omitempty"`
	UPC              int64             `bson:"upc,omitempty" json:"upc,omitempty"`
	SKU              int64             `bson:"sku,omitempty" json:"sku,omitempty"`
	Name             string            `bson:"name,omitempty" json:"name,omitempty"`
	Origin           string            `bson:"origin,omitempty" json:"origin,omitempty"`
	DeviceID         uuuid.UUID        `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Price            float64           `bson:"price,omitempty" json:"price,omitempty"`
	SalePrice        float64           `bson:"sale_price,omitempty" json:"sale_price,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Ethylene         float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	Version          int64             `bson:"version,omitempty" json:"version,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

type marshalFlash struct {
	ID               objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FlashID          string            `bson:"flash_id,omitempty" json:"flash_id,omitempty"`
	ItemID           string            `bson:"item_id,omitempty" json:"item_id,omitempty"`
	UPC              int64             `bson:"upc,omitempty" json:"upc,omitempty"`
	SKU              int64             `bson:"sku,omitempty" json:"sku,omitempty"`
	Name             string            `bson:"name,omitempty" json:"name,omitempty"`
	Origin           string            `bson:"origin,omitempty" json:"origin,omitempty"`
	DeviceID         string            `bson:"device_id,omitempty" json:"device_id,omitempty"`
	Price            float64           `bson:"price,omitempty" json:"price,omitempty"`
	SalePrice        float64           `bson:"sale_price,omitempty" json:"sale_price,omitempty"`
	Timestamp        int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Ethylene         float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	Version          int64             `bson:"version,omitempty" json:"version,omitempty"`
	AggregateVersion int64             `bson:"aggregate_version,omitempty" json:"aggregate_version,omitempty"`
}

func (f Flash) MarshalJSON() ([]byte, error) {
	mf := &marshalFlash{
		ID:               f.ID,
		UPC:              f.UPC,
		SKU:              f.SKU,
		Name:             f.Name,
		Origin:           f.Origin,
		Price:            f.Price,
		SalePrice:        f.SalePrice,
		Timestamp:        f.Timestamp,
		Ethylene:         f.Ethylene,
		Version:          f.Version,
		AggregateVersion: f.AggregateVersion,
	}

	if f.FlashID.String() != (uuuid.UUID{}).String() {
		mf.FlashID = f.FlashID.String()
	}

	if f.ItemID.String() != (uuuid.UUID{}).String() {
		mf.ItemID = f.ItemID.String()
	}

	if f.DeviceID.String() != (uuuid.UUID{}).String() {
		mf.DeviceID = f.DeviceID.String()
	}
	return json.Marshal(mf)
}

func (f *Flash) UnmarshalJSON(in []byte) error {
	var ok bool

	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	if m["_id"] != nil {
		f.ID = m["_id"].(objectid.ObjectID)
	}

	if m["flash_id"] != nil {
		f.FlashID, err = uuuid.FromString(m["flash_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["item_id"] != nil {
		f.ItemID, err = uuuid.FromString(m["item_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing ItemID for inventory")
			return err
		}
	}

	if m["device_id"] != nil {
		f.DeviceID, err = uuuid.FromString(m["device_id"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error parsing DeviceID for inventory")
			return err
		}
	}

	if m["upc"] != nil {
		upcType := reflect.TypeOf(m["upc"]).Kind()
		f.UPC, ok = m["upc"].(int64)
		if !ok {
			if upcType == reflect.Float64 {
				f.UPC = int64(m["upc"].(float64))
			} else {
				val, _ := strconv.Atoi((m["upc"]).(string))
				f.UPC = int64(val)
			}

		}
	}

	if m["sku"] != nil {
		skuType := reflect.TypeOf(m["sku"]).Kind()
		f.SKU, ok = m["sku"].(int64)
		if !ok {
			if skuType == reflect.Float64 {
				f.SKU = int64(m["sku"].(float64))
			} else {
				val, _ := strconv.Atoi((m["sku"]).(string))
				f.SKU = int64(val)
			}

		}
	}

	if m["price"] != nil {
		priceType := reflect.TypeOf(m["price"]).Kind()
		f.Price, ok = m["price"].(float64)
		if !ok {
			if priceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["price"]).(string))
				f.Price = float64(val)
			}
		}
	}

	if m["sale_price"] != nil {
		salePriceType := reflect.TypeOf(m["sale_price"]).Kind()
		f.SalePrice, ok = m["sale_price"].(float64)
		if !ok {
			if salePriceType != reflect.Float64 {
				val, _ := strconv.Atoi((m["sale_price"]).(string))
				f.SalePrice = float64(val)
			}
		}
	}

	if m["ethylene"] != nil {
		ethyleneType := reflect.TypeOf(m["ethylene"]).Kind()
		f.Ethylene, ok = m["ethylene"].(float64)
		if !ok {
			if ethyleneType != reflect.Float64 {
				val, _ := strconv.Atoi((m["ethylene"]).(string))
				f.Ethylene = float64(val)
			}
		}
	}

	if m["timestamp"] != nil {
		timestampType := reflect.TypeOf(m["timestamp"]).Kind()
		f.Timestamp, ok = m["timestamp"].(int64)
		if !ok {
			if timestampType == reflect.Float64 {
				f.Timestamp = int64(m["timestamp"].(float64))
			} else {
				val, _ := strconv.Atoi((m["timestamp"]).(string))
				f.Timestamp = int64(val)
			}
		}
	}

	if m["version"] != nil {
		versionType := reflect.TypeOf(m["version"]).Kind()
		f.Version, ok = m["version"].(int64)
		if !ok {
			if versionType == reflect.Float64 {
				f.Version = int64(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				f.Version = int64(val)
			}
		}
	}

	if m["aggregate_version"] != nil {
		aggregateVersionType := reflect.TypeOf(m["aggregate_version"]).Kind()
		f.AggregateVersion, ok = m["aggregate_version"].(int64)
		if !ok {
			if aggregateVersionType != reflect.Int64 {
				val, _ := strconv.Atoi((m["aggregate_version"]).(string))
				f.AggregateVersion = int64(val)
			}
		}
	}

	return nil
}
