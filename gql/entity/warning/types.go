package warning

import (
	"github.com/graphql-go/graphql"
)

// Inventory is the GraphQL-Type for Inventory model.
var Warning = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Warning",
		Fields: graphql.Fields{
			"warningID": &graphql.Field{
				Type: graphql.String,
			},
			"itemID": &graphql.Field{
				Type: graphql.String,
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
			"warningActive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ethylene": &graphql.Field{
				Type: graphql.Float,
			},
			"carbonDioxide": &graphql.Field{
				Type: graphql.Float,
			},
			"projectedDate": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var WarningInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "WarningInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"warningID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"itemID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
		"warningActive": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"ethylene": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"carbonDioxide": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"projectedDate": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var WarningUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WarningUpdateResult",
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

var WarningDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WarningDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
