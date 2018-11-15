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

	"Savings": &graphql.Field{
		Type:        graphql.NewList(SoldAvgTotal),
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
		Type:        graphql.NewList(SoldAvgTotal),
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

	"FlashSale": &graphql.Field{
		Type:        graphql.NewList(FlashSale),
		Description: "FlashSale",
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
}

var FlashSale = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FlashSaleType",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avg_sold": &graphql.Field{
				Type: graphql.Float,
			},
			"avg_total": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var SoldAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SoldAvgTotal",
		Fields: graphql.Fields{
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avg_sold": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var WasteAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WasteAvgTotal",
		Fields: graphql.Fields{
			"avg_total": &graphql.Field{
				Type: graphql.Float,
			},
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avg_waste": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DonateAvgTotal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DonateAvgTotal",
		Fields: graphql.Fields{
			"avg_total": &graphql.Field{
				Type: graphql.Float,
			},
			"_id": &graphql.Field{
				Type: DonateID,
			},
			"avg_donate": &graphql.Field{
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
