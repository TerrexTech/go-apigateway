package shipment

import (
	"github.com/TerrexTech/go-apigateway/gql"
	mt "github.com/TerrexTech/go-apigateway/gql/entity/shipment/mutationtype"
	"github.com/TerrexTech/go-apigateway/gql/entity/shipment/resolver"
	"github.com/graphql-go/graphql"
)

var mutationResolvers = map[string]graphql.FieldResolveFn{
	"addItem":    resolver.AddItem,
	"updateItem": resolver.UpdateItem,
	"deleteItem": resolver.DeleteItem,
}

// Mutations are GraphQL mutations for Auth.
var Mutations = map[string]*graphql.Field{
	"Shipment": &graphql.Field{
		Type:        mt.MutationsOutput,
		Description: "Shipment Operations",
		Args: graphql.FieldConfigArgument{
			"addItem": &graphql.ArgumentConfig{
				Type: mt.ItemInput,
			},
			"updateItem": &graphql.ArgumentConfig{
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: "ShipmentUpdateParams",
						Fields: graphql.InputObjectConfigFieldMap{
							"filter": &graphql.InputObjectFieldConfig{
								Type: mt.ItemInput,
							},
							"update": &graphql.InputObjectFieldConfig{
								Type: mt.ItemInput,
							},
						},
					},
				),
			},
			"deleteItem": &graphql.ArgumentConfig{
				Type: mt.ItemInput,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			result, err := gql.FormatResolvers(p, mutationResolvers)
			return result, err
		},
	},
}
