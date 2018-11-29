package types

import (
	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/graphql-go/graphql"
)

var ItemOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemOutput",
		Fields: graphql.Fields{
			"addItem": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "AddItem",
						Fields: graphql.Fields{
							"itemID": &graphql.Field{
								Type: graphql.String,
							},
							"dateArrived": &graphql.Field{
								Type: graphql.String,
							},
							"lot": &graphql.Field{
								Type: graphql.String,
							},
							"name": &graphql.Field{
								Type: graphql.String,
							},
							"origin": &graphql.Field{
								Type: graphql.String,
							},
							"price": &graphql.Field{
								Type: graphql.String,
							},
							"rsCustomerID": &graphql.Field{
								Type: graphql.String,
							},
							"sku": &graphql.Field{
								Type: graphql.String,
							},
							"timestamp": &graphql.Field{
								Type: graphql.String,
							},
							"totalWeight": &graphql.Field{
								Type: graphql.String,
							},
							"upc": &graphql.Field{
								Type: graphql.String,
							},
							"error": &graphql.Field{
								Type: gql.ResultError,
							},
						},
					},
				),
			},
		},
	},
)
