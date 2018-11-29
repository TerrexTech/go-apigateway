package schema

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/auth"
	"github.com/TerrexTech/go-apigateway/gql/entity/flashsale"
	"github.com/TerrexTech/go-apigateway/gql/entity/inventory"
	"github.com/TerrexTech/go-apigateway/gql/entity/report"
	"github.com/TerrexTech/go-apigateway/gql/entity/warning"
	"github.com/TerrexTech/go-apigateway/gql/entity/disposal"
	"github.com/TerrexTech/go-apigateway/gql/entity/donation"
	"github.com/TerrexTech/go-apigateway/gql/entity/device"
	"github.com/graphql-go/graphql"
)

// RootQuery is the schema-definition for GraphQL queries.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: createFields(
		auth.Queries,
		inventory.Queries,
		// sale.Queries,
		report.Queries,
		warning.Queries,
		flashsale.Queries,
		disposal.Queries,
		donation.Queries,
		device.Queries,
	),
})
