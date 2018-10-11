package gqltype

import "github.com/graphql-go/graphql"

var SearchReportResponses = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchReportResponse",
		Fields: graphql.Fields{
			"inventory": &graphql.Field{
				Type: graphql.String,
			},
			"metric": &graphql.Field{
				Type: graphql.String,
			},
			// "reporttype": &graphql.Field{
			// 	Type: graphql.String,
			// },
		},
	},
)
