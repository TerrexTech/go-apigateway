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

var FlashSale = func(params graphql.ResolveParams) (interface{}, error) {
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
			"avg_sold": {
				"$avg": "$flashSaleWeight",
			},
			"avg_total": {
				"$avg": "$totalWeight",
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
	log.Println(err)
	log.Println(aggResult)
	return aggResult, nil
}
