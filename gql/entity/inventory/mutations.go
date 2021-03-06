package inventory

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"InventoryInsert": &graphql.Field{
		Type:        Inventory,
		Description: "Inserts item into Inventory",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dateArrived": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"dateSold": &graphql.ArgumentConfig{
				Type: graphql.Int,
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
			"flashSaleWeight": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"soldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"timestamp": &graphql.ArgumentConfig{
				Type: graphql.Int,
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
			"onFlashSale": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"projectedDate": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.Insert,
	},

	"InventoryUpdate": &graphql.Field{
		Type:        InventoryUpdateResult,
		Description: "Update items from Inventory",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: InventoryInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: InventoryInput,
			},
		},
		Resolve: resolver.Update,
	},

	"InventoryDelete": &graphql.Field{
		Type:        InventoryDeleteResult,
		Description: "Delete items from Inventory",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"deviceID": &graphql.ArgumentConfig{
				Type: graphql.String,
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
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"upc": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Delete,
	},
}
