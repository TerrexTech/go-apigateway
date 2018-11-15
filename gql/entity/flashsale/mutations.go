package flashsale

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/flashsale/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"FlashSaleInsert": &graphql.Field{
		Type:        Sale,
		Description: "Inserts item into Inventory",
		Args: graphql.FieldConfigArgument{
			"flashSaleID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"items": &graphql.ArgumentConfig{
				Type: graphql.NewList(SaleInput),
			},
			"timestamp": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Insert,
	},
}
