package auth

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/auth/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Auth
var Queries = map[string]*graphql.Field{
	"authLogin": &graphql.Field{
		Type:        AuthResponse,
		Description: "User Login",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"userName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Login,
	},

	"authAccessToken": &graphql.Field{
		Type:        AuthResponse,
		Description: "User Authentication",
		Args: graphql.FieldConfigArgument{
			"refreshToken": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sub": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.AccessToken,
	},
}
