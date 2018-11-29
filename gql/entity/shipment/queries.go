package shipment

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/shipment/mutationtype"
	"github.com/TerrexTech/go-apigateway/gql"
	qt "github.com/TerrexTech/go-apigateway/gql/entity/shipment/querytype"
	"github.com/TerrexTech/go-apigateway/gql/entity/shipment/resolver"
	"github.com/graphql-go/graphql"
)

var queryResolvers = map[string]graphql.FieldResolveFn{
	"latest": resolver.Latest,
}

// Queries are GraphQL queries for Auth
var Queries = map[string]*graphql.Field{
	"Shipment": &graphql.Field{
		Type:        qt.QueriesOutput,
		Description: "Latest Items",
		Args: graphql.FieldConfigArgument{
			"latest": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutationtype.ItemInput),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			result, err := gql.FormatResolvers(p, queryResolvers)
			return result, err
		},
	},
}
