package donation

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/donation/resolver"
	"github.com/graphql-go/graphql"
)

// Mutations are GraphQL mutations for Inventory.
var Mutations = map[string]*graphql.Field{
	"DonationInsert": &graphql.Field{
		Type:        Donation,
		Description: "Inserts item into Donation",
		Args: graphql.FieldConfigArgument{
			"donationID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"donateWeight": &graphql.ArgumentConfig{
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

	"DonationUpdate": &graphql.Field{
		Type:        DonationUpdateResult,
		Description: "Update items from Donation",
		Args: graphql.FieldConfigArgument{
			"filter": &graphql.ArgumentConfig{
				Type: DonationInput,
			},
			"update": &graphql.ArgumentConfig{
				Type: DonationInput,
			},
		},
		Resolve: resolver.Update,
	},

	"DonationDelete": &graphql.Field{
		Type:        DonationDeleteResult,
		Description: "Delete items from Donation",
		Args: graphql.FieldConfigArgument{
			"itemID": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"donationID": &graphql.ArgumentConfig{
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
