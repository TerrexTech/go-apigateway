package flashsale

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/flashsale/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"FlashsaleInsert": &graphql.Field{
		Type:        Flashsale,
		Description: "Inserts item into Flashsale",
		Args: graphql.FieldConfigArgument{
			"flashsaleID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lot": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"soldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"timestamp": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"totalWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"unsoldWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"onFlashsale": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"projectedDate": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.Insert,
	},

	"FlashsaleUpdate": &graphql.Field{
		Type:        FlashsaleUpdateResult,
		Description: "Update items from Flashsale",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: FlashsaleInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: FlashsaleInput,
			},
		},
		Resolve: resolver.Update,
	},

	"FlashsaleDelete": &graphql.Field{
		Type:        FlashsaleDeleteResult,
		Description: "Delete items from Flashsale",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"flashsaleID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"lot": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sku": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Delete,
	},
}
