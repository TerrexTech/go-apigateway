package disposal

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/disposal/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"DisposalInsert": &graphql.Field{
		Type:        Disposal,
		Description: "Inserts item into Disposal",
		Args: graphql.FieldConfigArgument{
			"disposalID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"disposalWeight": &graphql.ArgumentConfig{
				Type: graphql.Float,
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
		},
		Resolve: resolver.Insert,
	},

	"DisposalUpdate": &graphql.Field{
		Type:        DisposalUpdateResult,
		Description: "Update items from Disposal",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: DisposalInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: DisposalInput,
			},
		},
		Resolve: resolver.Update,
	},

	"DisposalDelete": &graphql.Field{
		Type:        DisposalDeleteResult,
		Description: "Delete items from Disposal",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"disposalID": &graphql.ArgumentConfig{
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
