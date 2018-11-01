package auth

import "github.com/graphql-go/graphql"

// AccessToken is the GraphQL-Type for AccessToken model.
var AccessToken = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AccessToken",
		Fields: graphql.Fields{
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// AuthResponse is the GraphQL-Type for AuthResponse model.
var AuthResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthResponse",
		Fields: graphql.Fields{
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
			"refreshToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// User is the GraphQL-Type for User model.
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"userName": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
