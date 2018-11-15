package report

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/report/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Auth
var Queries = map[string]*graphql.Field{
	"Donate": &graphql.Field{
		Type:        graphql.NewList(DonateAvgTotal),
		Description: "DonateAvgTotal",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Donate,
	},
}

var DonateAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonateAvgTotal",
		Fields: graphql.Fields{
			"avg_total": &graphql.Field{
				Type: graphql.Float,
			},
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avg_donate": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DonateID = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonateID",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
