package device

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/device/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Device
var Queries = map[string]*graphql.Field{
	"DeviceQueryItem": &graphql.Field{
		Type:        graphql.NewList(Device),
		Description: "Device Query",
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
		Resolve: resolver.QueryItem,
	},
	"DeviceQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Device),
		Description: "Device Query by Timestamp",
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
	"DeviceQueryCount": &graphql.Field{
		Type:        graphql.NewList(Device),
		Description: "Returns latest device items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}
