package types

import (
	"github.com/graphql-go/graphql"
)

var ItemInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Item",
		Fields: graphql.InputObjectConfigFieldMap{
			"itemID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dateArrived": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"lot": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"origin": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"rsCustomerID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"sku": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"timestamp": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"totalWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"upc": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
