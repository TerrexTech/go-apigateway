package gqltype

import "github.com/graphql-go/graphql"

var SearchReportResponses = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchReportResponse",
		Fields: graphql.Fields{
			"inventory": &graphql.Field{
				Type: Inventory,
			},
			"metric": &graphql.Field{
				Type: Metric,
			},
			// "reporttype": &graphql.Field{
			// 	Type: graphql.String,
			// },
		},
	},
)
