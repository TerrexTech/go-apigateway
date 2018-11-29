package warning

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/warning/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Warning
var Queries = map[string]*graphql.Field{
	"WarningQueryItem": &graphql.Field{
		Type:        graphql.NewList(Warning),
		Description: "Warning Query",
		Args: graphql.FieldConfigArgument{
			"warningID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
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
			"soldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"timestamp": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"totalWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"unsoldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"warningActive": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"ethylene": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"carbonDioxide": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"projectedDate": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.QueryItem,
	},
	"WarningQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Warning),
		Description: "Warning Query by Timestamp",
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
	"WarningQueryCount": &graphql.Field{
		Type:        graphql.NewList(Warning),
		Description: "Returns latest warning items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}
