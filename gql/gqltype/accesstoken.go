package gqltype

import "github.com/graphql-go/graphql"

// AccessToken represents the GraphQL-type for AccessToken model.
var AccessToken = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AccessToken",
		Fields: graphql.Fields{
			"access_token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
