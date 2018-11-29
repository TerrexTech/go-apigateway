package mutationtype

import (
	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/graphql-go/graphql"
)

var MutationsOutput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MutationsOutput",
		Fields: graphql.Fields{
			"register": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "Register",
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
		},
	},
)
