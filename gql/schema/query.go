package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/auth"
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory"
	"github.com/graphql-go/graphql"
)

// RootQuery is the schema-definition for GraphQL queries.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: createFields(
		auth.Queries,
		inventory.Queries,
		// sale.Queries,
	),
})
