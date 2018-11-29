package querytype

import "github.com/graphql-go/graphql"

var Credentials = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Credentials",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
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
