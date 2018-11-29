package querytype

import (
	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/TerrexTech/go-apigateway/gql/entity/shipment/mutationtype"
	"github.com/graphql-go/graphql"
)

// QueriesOutput is the GraphQL-Type for QueriesOutput model.
var QueriesOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "QueriesOutput",
		Fields: graphql.Fields{
			"latest": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "Latest",
						Fields: graphql.Fields{
							"items": &graphql.Field{
								Type: graphql.NewList(mutationtype.ItemOutput),
							},
							"error": &graphql.Field{
								Type: gql.ResultError,
							},
						},
					},
				),
			},

			"refreshToken": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "RefreshToken",
						Fields: graphql.Fields{
							"accessToken": &graphql.Field{
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
