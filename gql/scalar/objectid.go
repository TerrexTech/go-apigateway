package scalar

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

var ObjectID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ObjectID",
	Description: "MongoDB ObjectID Type",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {
		switch t := value.(type) {
		case objectid.ObjectID:
			return t.Hex()
		default:
			return objectid.NilObjectID
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch t := value.(type) {
		case string:
			id, err := objectid.FromHex(t)
			if err != nil {
				return objectid.NilObjectID
			}
			return id
		default:
			return objectid.NilObjectID
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		value := valueAST.GetValue()
		switch t := value.(type) {
		case string:
			id, err := objectid.FromHex(t)
			if err != nil {
				return objectid.NilObjectID
			}
			return id
		default:
			return objectid.NilObjectID
		}
	},
})
