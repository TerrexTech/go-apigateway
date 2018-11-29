package donation

import (
	"github.com/graphql-go/graphql"
)

// Inventory is the GraphQL-Type for Inventory model.
var Donation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Donation",
		Fields: graphql.Fields{
			"donationID": &graphql.Field{
				Type: graphql.String,
			},
			"itemID": &graphql.Field{
				Type: graphql.String,
			},
			"donateWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"lot": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"soldWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"timestamp": &graphql.Field{
				Type: graphql.Int,
			},
			"totalWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"unsoldWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DonationInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DonationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"donationID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"itemID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"donateWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"lot": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sku": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"soldWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"timestamp": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"totalWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"unsoldWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var DonationUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonationUpdateResult",
		Fields: graphql.Fields{
			"matchedCount": &graphql.Field{
				Type: graphql.Float,
			},
			"modifiedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DonationDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonationDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
