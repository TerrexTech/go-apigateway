package mutationtype

import (
	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/graphql-go/graphql"
)

var MutationsOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MutationsOutputShipment",
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

			"updateItem": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "UpdateItem",
						Fields: graphql.Fields{
							"filter": &graphql.Field{
								Type: ItemOutput,
							},
							"update": &graphql.Field{
								Type: ItemOutput,
							},
							"error": &graphql.Field{
								Type: gql.ResultError,
							},
						},
					},
				),
			},

			"deleteItem": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "DeleteItem",
						Fields: graphql.Fields{
							"matchedCount": &graphql.Field{
								Type: graphql.Int,
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
