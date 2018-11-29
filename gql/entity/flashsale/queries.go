package flashsale

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/flashsale/resolver"
	"github.com/graphql-go/graphql"
)

var Queries = map[string]*graphql.Field{
	"FlashsaleQueryItem": &graphql.Field{
		Type:        graphql.NewList(Flashsale),
		Description: "Flashsale Query",
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
		Resolve: resolver.QueryItem,
	},
	"FlashsaleQueryTimestamp": &graphql.Field{
		Type:        graphql.NewList(Flashsale),
		Description: "Flashsale Query by Timestamp",
		Args: graphql.FieldConfigArgument{
			"start": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"end": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryTimestamp,
	},
	"FlashsaleQueryCount": &graphql.Field{
		Type:        graphql.NewList(Flashsale),
		Description: "Returns latest flashsale items as specified in count",
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.QueryCount,
	},
}

// // Queries are GraphQL queries for Inventory
// var Queries = map[string]*graphql.Field{
// 	// "InventoryQueryItem": &graphql.Field{
// 	// 	Type:        graphql.NewList(Inventory),
// 	// 	Description: "Inventory Query",
// 	// 	Args: graphql.FieldConfigArgument{
// 	// 		"itemID": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"dateArrived": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"dateSold": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"deviceID": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"donateWeight": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"lot": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"name": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"origin": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"price": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"rsCustomerID": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"salePrice": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"sku": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"soldWeight": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"timestamp": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"totalWeight": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 		"upc": &graphql.ArgumentConfig{
// 	// 			Type: graphql.String,
// 	// 		},
// 	// 		"wasteWeight": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Float,
// 	// 		},
// 	// 	},
// 	// 	Resolve: resolver.QueryItem,
// 	// },
// 	// "InventoryQueryTimestamp": &graphql.Field{
// 	// 	Type:        graphql.NewList(Inventory),
// 	// 	Description: "Inventory Query by Timestamp",
// 	// 	Args: graphql.FieldConfigArgument{
// 	// 		"start": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Int,
// 	// 		},
// 	// 		"end": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Int,
// 	// 		},
// 	// 		"count": &graphql.ArgumentConfig{
// 	// 			Type: graphql.Int,
// 	// 		},
// 	// 	},
// 	// 	Resolve: resolver.QueryTimestamp,
// 	// },
// 	"FlashSaleQueryCount": &graphql.Field{
// 		Type:        graphql.NewList(Sale),
// 		Description: "Returns latest sales as specified in count",
// 		Args: graphql.FieldConfigArgument{
// 			"count": &graphql.ArgumentConfig{
// 				Type: graphql.Int,
// 			},
// 		},
// 		Resolve: resolver.QueryCount,
// 	},
// }
