package donation

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/donation/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Donation
var Queries = map[string]*graphql.Field{
	"DonationQueryItem": &graphql.Field{
		Type:        graphql.NewList(Donation),
		Description: "Donation Query",
		Args: graphql.FieldConfigArgument{
			"donationID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
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
	"DonationQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Donation),
		Description: "Donation Query by Timestamp",
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
	"DonationQueryCount": &graphql.Field{
		Type:        graphql.NewList(Donation),
		Description: "Returns latest donation items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}
