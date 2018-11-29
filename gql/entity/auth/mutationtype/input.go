package mutationtype

import "github.com/graphql-go/graphql"

var User = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "User",
		Fields: graphql.InputObjectConfigFieldMap{
			"firstName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"lastName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"userName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"role": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
