package inventory

import (
	"github.com/graphql-go/graphql"
)

// Inventory is the GraphQL-Type for Inventory model.
var Inventory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Inventory",
		Fields: graphql.Fields{
			"itemID": &graphql.Field{
				Type: graphql.String,
			},
			"dateArrived": &graphql.Field{
				Type: graphql.Int,
			},
			"dateSold": &graphql.Field{
				Type: graphql.Int,
			},
			"deviceID": &graphql.Field{
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
			"origin": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"rsCustomerID": &graphql.Field{
				Type: graphql.String,
			},
			"flashSaleWeight": &graphql.Field{
				Type: graphql.Boolean,
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
			"upc": &graphql.Field{
				Type: graphql.String,
			},
			"wasteWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"onFlashSale": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var InventoryInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InventoryInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"itemID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"dateArrived": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"dateSold": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"deviceID": &graphql.InputObjectFieldConfig{
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
		"origin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"rsCustomerID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"flashSaleWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
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
		"upc": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"wasteWeight": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"onFlashSale": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})

var InventoryUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InventoryUpdateResult",
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

var InventoryDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InventoryDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
