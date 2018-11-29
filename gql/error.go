package gql

import "github.com/graphql-go/graphql"

type Error struct {
	Code    int16  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var ResultError = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResultError",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
