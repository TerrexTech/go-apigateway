package disposal

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/disposal/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Disposal
var Queries = map[string]*graphql.Field{
	"DisposalQueryItem": &graphql.Field{
		Type:        graphql.NewList(Disposal),
		Description: "Disposal Query",
		Args: graphql.FieldConfigArgument{
			"disposalID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"disposalWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
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
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.QueryItem,
	},
	"DisposalQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Disposal),
		Description: "Disposal Query by Timestamp",
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
	"DisposalQueryCount": &graphql.Field{
		Type:        graphql.NewList(Disposal),
		Description: "Returns latest disposal items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}
