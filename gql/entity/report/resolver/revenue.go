package resolver

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
)

type RevenueResult struct {
	SKU            string  `bson:"sku,omitempty" json:"sku,omitempty"`
	Name           string  `bson:"name,omitempty" json:"name,omitempty"`
	PrevSoldWeight float64 `bson:"prevSoldWeight,omitempty" json:"prevSoldWeight,omitempty"`
	SoldWeight     float64 `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	// TotalWeight    float64 `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	RevenuePrev    float64 `bson:"revenuePrev,omitempty" json:"revenuePrev,omitempty"`
	RevenueCurr    float64 `bson:"revenueCurr,omitempty" json:"revenueCurr,omitempty"`
	RevenuePercent float64 `bson:"revenuePercent,omitempty" json:"revenuePercent,omitempty"`
}

var Revenue = func(params graphql.ResolveParams) (interface{}, error) {
	var reportAgg []RevenueResult
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
			"avgSold": {
				"$avg": "$soldWeight",
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
		currSoldWeight := m["avgSold"].(float64)
		prevSoldWeight := currSoldWeight / GenFloat(0.1, 2.8)

		revenueCurrRandPrice := GenFloat(0.5, 3.4)
		revenueCurr := currSoldWeight * revenueCurrRandPrice

		revenuePrev := prevSoldWeight * GenFloat(0.1, revenueCurrRandPrice)
		revenuePercent := ((revenueCurr - revenuePrev) / (revenuePrev * 4)) * 100

		reportAgg = append(reportAgg, RevenueResult{
			SKU:            sku,
			Name:           name,
			SoldWeight:     currSoldWeight,
			PrevSoldWeight: prevSoldWeight,
			RevenuePrev:    revenuePrev,
			RevenueCurr:    revenueCurr,
			RevenuePercent: revenuePercent,
		})
	}
	log.Println(err)
	log.Println(reportAgg)
	return reportAgg, nil
}
