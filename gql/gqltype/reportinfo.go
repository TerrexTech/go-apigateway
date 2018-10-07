package gqltype

import "github.com/graphql-go/graphql"

// AuthResponse represents the GraphQL-type for AuthResponse model.
var ReportResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ReportResponse",
		Fields: graphql.Fields{
			"report_id": &graphql.Field{
				Type: graphql.String,
			},
			"timestamp": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// var InventoryReponse = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "InventoryResponse",
// 		Fields: graphql.Fields{
// 			"report_id": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"report_type": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 		},
// 	},
// )
