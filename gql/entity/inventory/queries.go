package inventory

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Inventory
var Queries = map[string]*graphql.Field{
	"InventoryQuery": &graphql.Field{
		Type:        graphql.NewList(Inventory),
		Description: "Inventory Query",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"barcode": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dateArrived": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"dateSold": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"deviceID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"donateWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"expiryDate": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"lot": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"origin": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"rsCustomerID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"salePrice": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"soldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"timestamp": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"totalWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"upc": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"wasteWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Query,
	},
}
