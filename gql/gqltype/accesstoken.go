package gqltype

import "github.com/graphql-go/graphql"

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
