package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/gqltype"
	resolver "github.com/TerrexTech/go-apigateway/gql/resolver"
	"github.com/graphql-go/graphql"
)

// RootQuery is the schema-definition for GraphQL queries.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"login": &graphql.Field{
			Type:        gqltype.AuthResponse,
			Description: "User Authentication",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.LoginResolver,
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
		"createData": &graphql.Field{
			Type:        gqltype.ReportResponse,
			Description: "Create data",
			Args: graphql.FieldConfigArgument{
				"test": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// 	"timestamp": &graphql.ArgumentConfig{
				// 		Type: graphql.DateTime,
				// 	},
				// 	"report_id": &graphql.ArgumentConfig{
				// 		Type: graphql.String,
				// 	},
			},
			Resolve: resolver.CreateReportData,
		},
	},
})
