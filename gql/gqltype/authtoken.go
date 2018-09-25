package gqltype

import "github.com/graphql-go/graphql"

// AuthResponse represents the GraphQL-type for AuthResponse model.
var AuthResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthResponse",
		Fields: graphql.Fields{
			"access_token": &graphql.Field{
				Type: graphql.String,
			},
			"refresh_token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
