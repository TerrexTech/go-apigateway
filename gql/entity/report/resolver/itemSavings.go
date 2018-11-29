package resolver

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
)

type ReportResult struct {
	SKU             string  `bson:"sku,omitempty" json:"sku,omitempty"`
	Name            string  `bson:"name,omitempty" json:"name,omitempty"`
	PrevWasteWeight float64 `bson:"prevWasteWeight,omitempty" json:"prevWasteWeight,omitempty"`
	WasteWeight     float64 `bson:"wasteWeight,omitempty" json:"wasteWeight,omitempty"`
	// TotalWeight float64 `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	AmWastePrev    float64 `bson:"amWastePrev,omitempty" json:"amWastePrev,omitempty"`
	AmWasteCurr    float64 `bson:"amWasteCurr,omitempty" json:"amWasteCurr,omitempty"`
	SavingsPercent float64 `bson:"savingsPercent,omitempty" json:"savingsPercent,omitempty"`
}

var Savings = func(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	coll := rootValue["inventoryColl"].(*mongo.Collection)

	if params.Args["$lt"] == 0 || params.Args["$gt"] == 0 {
		err := errors.New("Missing timestamp value")
		log.Println(err)
		return nil, err
	}

	params.Args["timestamp"] = map[string]interface{}{
		"$lt": params.Args["lt"],
		"$gt": params.Args["gt"],
	}
	delete(params.Args, "lt")
	delete(params.Args, "gt")

	input, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "Unable to marshal params.Args")
		log.Println(err)
		return nil, err
	}

	log.Println(string(input))
	log.Println(params.Args)

	pipelineBuilder := fmt.Sprintf(`[
		{
			"$match": %s
		},
		{
			"$group" : {
			"_id" : {"sku" : "$sku","name":"$name"},
			"avgWaste": {
				"$avg": "$wasteWeight",
			}
		}
		}
	]`, input)

	pipelineAgg, err := bson.ParseExtJSONArray(pipelineBuilder)
	if err != nil {
		err = errors.Wrap(err, "Query: Error in generating pipeline for report")
		log.Println(err)
		return nil, err
	}

	aggResult, err := coll.Aggregate(pipelineAgg)
	if err != nil {
		err = errors.Wrap(err, "Query: Error in getting aggregate results ")
		log.Println(err)
		return nil, err
	}

	var reportAgg []ReportResult

	for _, v := range aggResult {
		m, assertOK := v.(map[string]interface{})
		if !assertOK {
			err := errors.New("Error getting results ")
			log.Println(err)
		}

		groupByFields := m["_id"]
		mapInGroupBy := groupByFields.(map[string]interface{})
		sku := mapInGroupBy["sku"].(string)
		name := mapInGroupBy["name"].(string)

		//Generate value for previous year
		currWasteWeight := m["avgWaste"].(float64)
		prevWasteWeight := currWasteWeight * GenFloat(1.2, 3.4)

		amWasteCurrRandPrice := GenFloat(0.5, 3.4)
		amWasteCurr := currWasteWeight * amWasteCurrRandPrice

		amWastePrev := prevWasteWeight * GenFloat(0.1, amWasteCurrRandPrice)
		savingsPercent := ((amWastePrev - amWasteCurr) / (amWastePrev)) * 100

		reportAgg = append(reportAgg, ReportResult{
			SKU:             sku,
			Name:            name,
			WasteWeight:     currWasteWeight,
			PrevWasteWeight: prevWasteWeight,
			AmWastePrev:     amWastePrev,
			AmWasteCurr:     amWasteCurr,
			SavingsPercent:  savingsPercent,
		})
	}
	log.Println(err)
	log.Println(reportAgg)
	return reportAgg, nil
}

func GenFloat(min float64, max float64) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random := min + r1.Float64()*(max-min)
	return random
}
