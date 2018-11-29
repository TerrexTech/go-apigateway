package querytype

import "github.com/graphql-go/graphql"

var Filters = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Filters",
		Fields: graphql.InputObjectConfigFieldMap{
			"latest": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"userName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
