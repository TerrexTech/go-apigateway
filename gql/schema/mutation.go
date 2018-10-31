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
			Type:        gqltype.SearchReportResponses,
			Description: "Report Generation",
			Args: graphql.FieldConfigArgument{
				"item_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rs_customer_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"report_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"report_type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"version": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.CreateReportData,
		},
		"addFlashSale": &graphql.Field{
			Type:        gqltype.SearchReportResponses,
			Description: "Flash sale insert",
			Args: graphql.FieldConfigArgument{
				"flash_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"item_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"upc": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"sku": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"origin": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.CreateReportData,
		},
	},
})
