package auth

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/auth/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Auth.
var Mutations = map[string]*graphql.Field{
	"authRegister": &graphql.Field{
		Type:        AuthResponse,
		Description: "Registers new User",
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
			"userName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"role": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Register,
	},
}
