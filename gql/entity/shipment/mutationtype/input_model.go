package mutationtype

import "github.com/graphql-go/graphql"

var ItemInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ItemInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"itemID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dateArrived": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
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
				Type: graphql.Float,
			},
			"rsCustomerID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"sku": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"timestamp": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"totalWeight": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"upc": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
