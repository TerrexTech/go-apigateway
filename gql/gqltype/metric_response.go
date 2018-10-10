package gqltype

import (
	"github.com/TerrexTech/go-apigateway/gql/scalar"
	"github.com/graphql-go/graphql"
)

var Metric = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Metric",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: scalar.ObjectID,
			},
			"item_id": &graphql.Field{
				Type: scalar.UUID,
			},
			"device_id": &graphql.Field{
				Type: scalar.UUID,
			},
			"timestamp": &graphql.Field{
				Type: graphql.Int,
			},
			"temp_in": &graphql.Field{
				Type: graphql.Float,
			},
			"humidity": &graphql.Field{
				Type: graphql.Float,
			},
			"ethylene": &graphql.Field{
				Type: graphql.Float,
			},
			"carbon_di": &graphql.Field{
				Type: graphql.Float,
			},
			"version": &graphql.Field{
				Type: graphql.Int,
			},
			"aggregate_id": &graphql.Field{
				Type: graphql.Int,
			},
			"aggregate_version": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
