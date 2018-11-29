package report

import (
	"github.com/TerrexTech/go-apigateway/gql/entity/report/resolver"
	"github.com/graphql-go/graphql"
)

// Queries are GraphQL queries for Auth
var Queries = map[string]*graphql.Field{
	"Donate": &graphql.Field{
		Type:        graphql.NewList(DonateAvgTotal),
		Description: "DonateAvgTotal",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Donate,
	},

	"Waste": &graphql.Field{
		Type:        graphql.NewList(WasteAvgTotal),
		Description: "WasteAvgTotal",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Waste,
	},

	"FlashSale": &graphql.Field{
		Type:        graphql.NewList(FlashSaleAvgTotal),
		Description: "FlashSale",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.FlashSale,
	},
	"ItemSold": &graphql.Field{
		Type:        graphql.NewList(ItemSoldAvgTotal),
		Description: "ItemSold",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.ItemSold,
	},

	"Savings": &graphql.Field{
		Type:        graphql.NewList(SavingAvgTotal),
		Description: "Savings",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Savings,
	},

	"Revenue": &graphql.Field{
		Type:        graphql.NewList(RevenueAvgTotal),
		Description: "Revenue",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.Revenue,
	},

	"EthyleneCO2": &graphql.Field{
		Type:        graphql.NewList(EthyleneAvgTotal),
		Description: "EthyleneCo2",
		Args: graphql.FieldConfigArgument{
			"lt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"gt": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: resolver.EthyleneCO2,
	},
}

var DonateAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonateAvgTotal",
		Fields: graphql.Fields{
			"avgTotal": &graphql.Field{
				Type: graphql.Float,
			},
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avgDonate": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var WasteAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WasteAvgTotal",
		Fields: graphql.Fields{
			"avgTotal": &graphql.Field{
				Type: graphql.Float,
			},
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avgWaste": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var FlashSaleAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FlashSaleType",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avgFlashSold": &graphql.Field{
				Type: graphql.Float,
			},
			"avgSold": &graphql.Field{
				Type: graphql.Float,
			},
			"avgTotal": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var ItemSoldAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemSoldType",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avgSold": &graphql.Field{
				Type: graphql.Float,
			},
			"avgTotal": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var SavingAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SavingsAvgTotal",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"wasteWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"prevWasteWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"amWastePrev": &graphql.Field{
				Type: graphql.Float,
			},
			"amWasteCurr": &graphql.Field{
				Type: graphql.Float,
			},
			"savingsPercent": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var RevenueAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RevenueAvgTotal",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"prevSoldWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"soldWeight": &graphql.Field{
				Type: graphql.Float,
			},
			"revenuePrev": &graphql.Field{
				Type: graphql.Float,
			},
			"revenueCurr": &graphql.Field{
				Type: graphql.Float,
			},
			"revenuePercent": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var EthyleneAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EthyleneAvgTotal",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: EthyleneID,
			},
			"avgEthylene": &graphql.Field{
				Type: graphql.Float,
			},
			"avgCarbonDioxide": &graphql.Field{
				Type: graphql.Float,
			},
			"avgTempIn": &graphql.Field{
				Type: graphql.Float,
			},
			"avgHumidity": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DonateID = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonateID",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var EthyleneID = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EthyleneID",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"sku": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
