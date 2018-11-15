package main

import (
	"encoding/json"

	util "github.com/TerrexTech/go-commonutils/commonutil"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// AggregateID is the global AggregateID for Inventory Aggregate.
const AggregateID int8 = 2

// Inventory defines the Inventory Aggregate.
type Inventory struct {
	ID              objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ItemID          uuuid.UUID        `bson:"itemID,omitempty" json:"itemID,omitempty"`
	DateArrived     int64             `bson:"dateArrived,omitempty" json:"dateArrived,omitempty"`
	DateSold        int64             `bson:"dateSold,omitempty" json:"dateSold,omitempty"`
	DeviceID        uuuid.UUID        `bson:"deviceID,omitempty" json:"deviceID,omitempty"`
	DonateWeight    float64           `bson:"donateWeight,omitempty" json:"donateWeight,omitempty"`
	Lot             string            `bson:"lot,omitempty" json:"lot,omitempty"`
	Name            string            `bson:"name,omitempty" json:"name,omitempty"`
	Origin          string            `bson:"origin,omitempty" json:"origin,omitempty"`
	Price           float64           `bson:"price,omitempty" json:"price,omitempty"`
	RSCustomerID    uuuid.UUID        `bson:"rsCustomerID,omitempty" json:"rsCustomerID,omitempty"`
	FlashSaleWeight float64           `bson:"flashSaleWeight,omitempty" json:"flashSaleWeight,omitempty"`
	SKU             string            `bson:"sku,omitempty" json:"sku,omitempty"`
	SoldWeight      float64           `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	Timestamp       int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	TotalWeight     float64           `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	UPC             string            `bson:"upc,omitempty" json:"upc,omitempty"`
	WasteWeight     float64           `bson:"wasteWeight,omitempty" json:"wasteWeight,omitempty"`
	OnFlashSale     bool              `bson:"onFlashSale,omitempty" json:"onFlashSale,omitempty"`
}

// MarshalBSON returns bytes of BSON-type.
func (i Inventory) MarshalBSON() ([]byte, error) {
	in := map[string]interface{}{
		"itemID":          i.ItemID.String(),
		"dateArrived":     i.DateArrived,
		"dateSold":        i.DateSold,
		"deviceID":        i.DeviceID.String(),
		"donateWeight":    i.DonateWeight,
		"lot":             i.Lot,
		"name":            i.Name,
		"origin":          i.Origin,
		"onFlashSale":     i.OnFlashSale,
		"price":           i.Price,
		"rsCustomerID":    i.RSCustomerID.String(),
		"flashSaleWeight": i.FlashSaleWeight,
		"sku":             i.SKU,
		"soldWeight":      i.SoldWeight,
		"timestamp":       i.Timestamp,
		"totalWeight":     i.TotalWeight,
		"upc":             i.UPC,
		"wasteWeight":     i.WasteWeight,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID
	}
	return bson.Marshal(in)
}

// MarshalJSON returns bytes of JSON-type.
func (i *Inventory) MarshalJSON() ([]byte, error) {
	in := map[string]interface{}{
		"itemID":          i.ItemID.String(),
		"dateArrived":     i.DateArrived,
		"dateSold":        i.DateSold,
		"deviceID":        i.DeviceID.String(),
		"donateWeight":    i.DonateWeight,
		"lot":             i.Lot,
		"name":            i.Name,
		"origin":          i.Origin,
		"onFlashSale":     i.OnFlashSale,
		"price":           i.Price,
		"rsCustomerID":    i.RSCustomerID.String(),
		"flashSaleWeight": i.FlashSaleWeight,
		"sku":             i.SKU,
		"soldWeight":      i.SoldWeight,
		"timestamp":       i.Timestamp,
		"totalWeight":     i.TotalWeight,
		"upc":             i.UPC,
		"wasteWeight":     i.WasteWeight,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID.Hex()
	}
	return json.Marshal(in)
}

// UnmarshalBSON returns BSON-type from bytes.
func (i *Inventory) UnmarshalBSON(in []byte) error {
	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

// UnmarshalJSON returns JSON-type from bytes.
func (i *Inventory) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

// unmarshalFromMap unmarshals Map into Inventory.
func (i *Inventory) unmarshalFromMap(m map[string]interface{}) error {
	var err error
	var assertOK bool

	// Hoping to discover a better way to do this someday
	if m["_id"] != nil {
		i.ID, assertOK = m["_id"].(objectid.ObjectID)
		if !assertOK {
			i.ID, err = objectid.FromHex(m["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if m["itemID"] != nil {
		i.ItemID, err = uuuid.FromString(m["itemID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting ItemID")
			return err
		}
	}

	if m["deviceID"] != nil {
		i.DeviceID, err = uuuid.FromString(m["deviceID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DeviceID")
			return err
		}
	}

	if m["rsCustomerID"] != nil {
		i.RSCustomerID, err = uuuid.FromString(m["rsCustomerID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting RSCustomerID")
			return err
		}
	}

	if m["dateArrived"] != nil {
		i.DateArrived, err = util.AssertInt64(m["dateArrived"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DateArrived")
			return err
		}
	}
	if m["dateSold"] != nil {
		i.DateSold, err = util.AssertInt64(m["dateSold"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DateSold")
			return err
		}
	}
	if m["donateWeight"] != nil {
		i.DonateWeight, err = util.AssertFloat64(m["donateWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting DonateWeight")
			return err
		}
	}
	if m["lot"] != nil {
		i.Lot, assertOK = m["lot"].(string)
		if !assertOK {
			return errors.New("Error while asserting Lot")
		}
	}
	if m["name"] != nil {
		i.Name, assertOK = m["name"].(string)
		if !assertOK {
			return errors.New("Error while asserting Name")
		}
	}
	if m["origin"] != nil {
		i.Origin, assertOK = m["origin"].(string)
		if !assertOK {
			return errors.New("Error while asserting Origin")
		}
	}
	if m["price"] != nil {
		i.Price, err = util.AssertFloat64(m["price"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Price")
			return err
		}
	}
	if m["flashSaleWeight"] != nil {
		i.FlashSaleWeight, err = util.AssertFloat64(m["flashSaleWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting FlashSaleWeight")
			return err
		}
	}
	if m["sku"] != nil {
		i.SKU, assertOK = m["sku"].(string)
		if !assertOK {
			return errors.New("Error while asserting Sku")
		}
	}
	if m["soldWeight"] != nil {
		i.SoldWeight, err = util.AssertFloat64(m["soldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting SoldWeight")
			return err
		}
	}
	if m["timestamp"] != nil {
		i.Timestamp, err = util.AssertInt64(m["timestamp"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Timestamp")
			return err
		}
	}
	if m["totalWeight"] != nil {
		i.TotalWeight, err = util.AssertFloat64(m["totalWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting TotalWeight")
			return err
		}
	}
	if m["upc"] != nil {
		i.UPC, assertOK = m["upc"].(string)
		if !assertOK {
			return errors.New("Error while asserting UPC")
		}
	}
	if m["onFlashSale"] != nil {
		i.OnFlashSale, assertOK = m["onFlashSale"].(bool)
		if !assertOK {
			return errors.New("Error while asserting OnFlashSale")
		}
	}
	if m["wasteWeight"] != nil {
		i.WasteWeight, err = util.AssertFloat64(m["wasteWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting WasteWeight")
			return err
		}
	}

	return nil
}
