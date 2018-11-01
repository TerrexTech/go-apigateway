package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/auth"
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory"
	"github.com/graphql-go/graphql"
)

// RootMutation is the schema-definition for GraphQL mutations.
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: createFields(
		auth.Mutations,
		inventory.Mutations,
	),
})
