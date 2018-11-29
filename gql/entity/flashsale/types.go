package flashsale

import (
	"github.com/graphql-go/graphql"
)

var Flashsale = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Flashsale",
		Fields: graphql.Fields{
			"flashsaleID": &graphql.Field{
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
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"onFlashsale": &graphql.Field{
				Type: graphql.Boolean,
			},
			"projectedDate": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FlashsaleInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FlashsaleInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"flashsaleID": &graphql.InputObjectFieldConfig{
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
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"onFlashsale": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"projectedDate": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var FlashsaleUpdateResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FlashsaleUpdateResult",
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

var FlashsaleDeleteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FlashsaleDeleteResult",
		Fields: graphql.Fields{
			"deletedCount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// // Sale is the GraphQL-Type for Sale model.
// var Sale = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "FlashSale",
// 		Fields: graphql.Fields{
// 			"flashSaleID": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"items": &graphql.Field{
// 				Type: graphql.NewList(SaleInput),
// 			},
// 			"timestamp": &graphql.Field{
// 				Type: graphql.Float,
// 			},
// 		},
// 	},
// )

// var SaleInput = graphql.NewInputObject(
// 	graphql.InputObjectConfig{
// 		Name: "FlashSaleInput",
// 		Fields: graphql.InputObjectConfigFieldMap{
// 			"itemID": &graphql.InputObjectFieldConfig{
// 				Type: graphql.String,
// 			},
// 			"upc": &graphql.InputObjectFieldConfig{
// 				Type: graphql.String,
// 			},
// 			"weight": &graphql.InputObjectFieldConfig{
// 				Type: graphql.Float,
// 			},
// 			"lot": &graphql.InputObjectFieldConfig{
// 				Type: graphql.String,
// 			},
// 			"sku": &graphql.InputObjectFieldConfig{
// 				Type: graphql.String,
// 			},
// 		},
// 	},
// )
