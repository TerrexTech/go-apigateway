package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/gqltype"
	resolver "github.com/TerrexTech/go-apigateway/gql/resolver"
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"register": &graphql.Field{
			Type:        gqltype.AuthResponse,
			Description: "User Authentication",
			Args: graphql.FieldConfigArgument{
				"firstName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"lastName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"role": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.RegisterResolver,
		},

		"accessToken": &graphql.Field{
			Type:        gqltype.AuthResponse,
			Description: "User Authentication",
			Args: graphql.FieldConfigArgument{
				"refreshToken": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sub": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.AccessTokenResolver,
		},
	},
})
