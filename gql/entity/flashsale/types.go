package flashsale

import (
	"github.com/graphql-go/graphql"
)

// Sale is the GraphQL-Type for Sale model.
var Sale = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FlashSale",
		Fields: graphql.Fields{
			"flashSaleID": &graphql.Field{
				Type: graphql.String,
			},
			"items": &graphql.Field{
				Type: graphql.NewList(SaleInput),
			},
			"timestamp": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var SaleInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "FlashSaleInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"itemID": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"upc": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"weight": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"lot": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"sku": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
