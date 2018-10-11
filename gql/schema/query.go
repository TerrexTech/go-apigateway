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
			Type:        gqltype.SearchReportResponses,
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
		"searchReport": &graphql.Field{
			Type:        gqltype.SearchReportResponses,
			Description: "Generate report based on search",
			Args: graphql.FieldConfigArgument{
				"inventory": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"metric": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				// "reporttype": &graphql.ArgumentConfig{
				// 	Type: graphql.String,
				// },
			},
			Resolve: resolver.SearchReport,
		},
	},
})

// var searchFields = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "asd",
// 	Fields: graphql.Fields{
// 		"fieldname": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		// "datatype": &graphql.Field{
// 		// 	Type: graphql.String,
// 		// },
// 		// "equal": &graphql.Field{
// 		// 	Type: graphql.String,
// 		// },
// 		// "upper_limit": &graphql.Field{
// 		// 	Type: graphql.Float,
// 		// },
// 		// "lower_limit": &graphql.Field{
// 		// 	Type: graphql.Float,
// 		// },
// 	},
// })
