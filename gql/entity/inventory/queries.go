package inventory

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Inventory
var Queries = map[string]*graphql.Field{
	"InventoryQueryItem": &graphql.Field{
		Type:        graphql.NewList(Inventory),
		Description: "Inventory Query",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
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
				Type: graphql.String,
			},
			"wasteWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.QueryItem,
	},
	"InventoryQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Inventory),
		Description: "Inventory Query by Timestamp",
		Args: graphql.FieldConfigArgument{
			"start": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"end": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryTimestamp,
	},
	"InventoryQueryCount": &graphql.Field{
		Type:        graphql.NewList(Inventory),
		Description: "Returns latest inventory items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}
