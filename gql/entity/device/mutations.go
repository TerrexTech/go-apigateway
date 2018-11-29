package device

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/device/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"DeviceInsert": &graphql.Field{
		Type:        Device,
		Description: "Inserts item into Device",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"deviceID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"dateInstalled": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"lot": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lastMaintenance": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Insert,
	},

	"DeviceUpdate": &graphql.Field{
		Type:        DeviceUpdateResult,
		Description: "Update items from Device",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: DeviceInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: DeviceInput,
			},
		},
		Resolve: resolver.Update,
	},

	"DeviceDelete": &graphql.Field{
		Type:        DeviceDeleteResult,
		Description: "Delete items from Device",
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
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Delete,
	},
}
