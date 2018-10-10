package gqltype

import (
	"github.com/TerrexTech/go-apigateway/gql/scalar"
	"github.com/graphql-go/graphql"
)

var Inventory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Inventory",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: scalar.ObjectID,
			},
			"item_id": &graphql.Field{
				Type: scalar.UUID,
			},
			"upc": &graphql.Field{
				Type: graphql.Int,
			},
			"sku": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"origin": &graphql.Field{
				Type: graphql.String,
			},
			"device_id": &graphql.Field{
				Type: scalar.UUID,
			},
			"total_weight": &graphql.Field{
				Type: graphql.Float,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"date_arrived": &graphql.Field{
				Type: graphql.Int,
			},
			"expiry_date": &graphql.Field{
				Type: graphql.Int,
			},
			"timestamp": &graphql.Field{
				Type: graphql.Int,
			},
			"rs_customer_id": &graphql.Field{
				Type: scalar.UUID,
			},
			"waste_weight": &graphql.Field{
				Type: graphql.Float,
			},
			"donate_weight": &graphql.Field{
				Type: graphql.Float,
			},
			"aggregate_version": &graphql.Field{
				Type: graphql.Int,
			},
			"aggregate_id": &graphql.Field{
				Type: graphql.Int,
			},
			"date_sold": &graphql.Field{
				Type: graphql.Int,
			},
			"sale_price": &graphql.Field{
				Type: graphql.Float,
			},
			"sold_weight": &graphql.Field{
				Type: graphql.Float,
			},
			"prod_quantity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
