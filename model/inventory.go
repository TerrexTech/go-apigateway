package model

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

type Inventory struct {
	ID               objectid.ObjectID `json:"_id,omitempty"`
	ItemID           uuuid.UUID        `json:"item_id,omitempty"`
	UPC              int64             `json:"upc,omitempty"`
	SKU              int64             `json:"sku,omitempty"`
	Name             string            `json:"name,omitempty"`
	Origin           string            `json:"origin,omitempty"`
	DeviceID         uuuid.UUID        `json:"device_id,omitempty"`
	TotalWeight      float64           `json:"total_weight,omitempty"`
	Price            float64           `json:"price,omitempty"`
	Lot              string            `json:"lot,omitempty"`
	DateArrived      int64             `json:"date_arrived,omitempty"`
	ExpiryDate       int64             `json:"expiry_date,omitempty"`
	Timestamp        int64             `json:"timestamp,omitempty"`
	RsCustomerID     uuuid.UUID        `json:"rs_customer_id,omitempty"`
	WasteWeight      float64           `json:"waste_weight,omitempty"`
	DonateWeight     float64           `json:"donate_weight,omitempty"`
	AggregateVersion int64             `json:"aggregate_version,omitempty"`
	AggregateID      int8              `json:"aggregate_id,omitempty"`
	DateSold         int64             `json:"date_sold,omitempty"`
	SalePrice        float64           `json:"sale_price,omitempty"`
	SoldWeight       float64           `json:"sold_weight,omitempty"`
	ProdQuantity     int64             `json:"prod_quantity,omitempty"`
	Version          int64             `json:"version,omitempty"`
}

type marshalInventory struct {
	ID               objectid.ObjectID `json:"_id,omitempty"`
	ItemID           string            `json:"item_id,omitempty"`
	UPC              int64             `json:"upc,omitempty"`
	SKU              int64             `json:"sku,omitempty"`
	Name             string            `json:"name,omitempty"`
	Origin           string            `json:"origin,omitempty"`
	DeviceID         string            `json:"device_id,omitempty"`
	TotalWeight      float64           `json:"total_weight,omitempty"`
	Price            float64           `json:"price,omitempty"`
	Lot              string            `json:"lot,omitempty"`
	DateArrived      int64             `json:"date_arrived,omitempty"`
	ExpiryDate       int64             `json:"expiry_date,omitempty"`
	Timestamp        int64             `json:"timestamp,omitempty"`
	RsCustomerID     string            `json:"rs_customer_id,omitempty"`
	WasteWeight      float64           `json:"waste_weight,omitempty"`
	DonateWeight     float64           `json:"donate_weight,omitempty"`
	AggregateVersion int64             `json:"aggregate_version,omitempty"`
	AggregateID      int8              `json:"aggregate_id,omitempty"`
	DateSold         int64             `json:"date_sold,omitempty"`
	SalePrice        float64           `json:"sale_price,omitempty"`
	SoldWeight       float64           `json:"sold_weight,omitempty"`
	ProdQuantity     int64             `json:"prod_quantity,omitempty"`
	Version          int64             `json:"version,omitempty"`
}

func (i *Inventory) MarshalJSON() ([]byte, error) {
	in := &marshalInventory{
		UPC:              i.UPC,
		SKU:              i.SKU,
		Name:             i.Name,
		Origin:           i.Origin,
		TotalWeight:      i.TotalWeight,
		Price:            i.Price,
		Lot:              i.Lot,
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
		Version:          i.Version,
		ProdQuantity:     i.ProdQuantity,
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

func (i *Inventory) UnmarshalJSON(in []byte) error {
	var ok bool
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	// if m["_id"] != nil {
	// 	i.ID, err = objectid.FromHex(m["_id"].(string))
	// }

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

	if m["lot"] != nil {
		i.Lot = m["lot"].(string)
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

	if m["version"] != nil {
		versionType := reflect.TypeOf(m["version"]).Kind()
		i.Version, ok = m["version"].(int64)
		if !ok {
			if versionType == reflect.Float64 {
				i.Version = int64(m["version"].(float64))
			} else {
				val, _ := strconv.Atoi((m["version"]).(string))
				i.Version = int64(val)
			}
		}
	}

	if m["prod_quantity"] != nil {
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
