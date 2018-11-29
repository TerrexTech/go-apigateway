package querytype

import (
	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/graphql-go/graphql"
)

// QueriesOutput is the GraphQL-Type for QueriesOutput model.
var QueriesOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "QueriesOutput",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "Login",
						Fields: graphql.Fields{
							"accessToken": &graphql.Field{
								Type: graphql.String,
							},
							"refreshToken": &graphql.Field{
								Type: graphql.String,
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
