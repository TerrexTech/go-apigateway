package warning

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/warning/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"WarningInsert": &graphql.Field{
		Type:        Warning,
		Description: "Inserts item into Warning",
		Args: graphql.FieldConfigArgument{
			"warningID": &graphql.ArgumentConfig{
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
			"warningActive": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"ethylene": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"carbonDioxide": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"projectedDate": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolver.Insert,
	},

	"WarningUpdate": &graphql.Field{
		Type:        WarningUpdateResult,
		Description: "Update items from Warning",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: WarningInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: WarningInput,
			},
		},
		Resolve: resolver.Update,
	},

	"WarningDelete": &graphql.Field{
		Type:        WarningDeleteResult,
		Description: "Delete items from Warning",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"warningID": &graphql.ArgumentConfig{
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
