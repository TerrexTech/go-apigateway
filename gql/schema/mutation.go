package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/gqltype"
	resolver "github.com/TerrexTech/go-apigateway/gql/resolver"
	"github.com/graphql-go/graphql"
)

// RootMutation is the schema-definition for GraphQL mutations.
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

		"report": &graphql.Field{
			Type:        gqltype.AuthResponse,
			Description: "Report Generation",
			Args: graphql.FieldConfigArgument{
				"item_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rs_customer_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"device_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"temp_in": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"humidity": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"ethylene": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"carbon_di": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: resolver.ReportResolver,
		},
	},
})
