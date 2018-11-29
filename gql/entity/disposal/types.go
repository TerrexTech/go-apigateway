package disposal

import (
	"github.com/graphql-go/graphql"
)

// Inventory is the GraphQL-Type for Inventory model.
var Disposal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Disposal",
		Fields: graphql.Fields{
			"disposalID": &graphql.Field{
				Type: graphql.String,
			},
			"itemID": &graphql.Field{
				Type: graphql.String,
			},
			"disposalWeight": &graphql.Field{
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

var DisposalInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DisposalInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"disposalID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"itemID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"disposalWeight": &graphql.InputObjectFieldConfig{
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

var DisposalUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DisposalUpdateResult",
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

var DisposalDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DisposalDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
