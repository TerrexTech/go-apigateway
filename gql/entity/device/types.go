package device

import (
	"github.com/graphql-go/graphql"
)

// Inventory is the GraphQL-Type for Inventory model.
var Device = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.Fields{
			"itemID": &graphql.Field{
				Type: graphql.String,
			},
			"deviceID": &graphql.Field{
				Type: graphql.String,
			},
			"dateInstalled": &graphql.Field{
				Type: graphql.Int,
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
			"lastMaintenance": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DeviceInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DeviceInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"itemID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"deviceID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"dateInstalled": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
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
		"lastMaintenance": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var DeviceUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DeviceUpdateResult",
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

var DeviceDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DeviceDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
