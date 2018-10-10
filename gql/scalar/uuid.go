package scalar

import (
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

var UUID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "UUID",
	Description: "TerrexTech/uuuid Type",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {
		switch t := value.(type) {
		case uuuid.UUID:
			return t.String()
		default:
			return objectid.NilObjectID
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch t := value.(type) {
		case string:
			uuid, err := uuuid.FromString(t)
			if err != nil {
				return uuuid.UUID{}
			}
			return uuid
		case []byte:
			uuid, err := uuuid.FromBytes(t)
			if err != nil {
				return uuuid.UUID{}
			}
			return uuid
		default:
			return uuuid.UUID{}
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		value := valueAST.GetValue()
		switch t := value.(type) {
		case string:
			uuid, err := uuuid.FromString(t)
			if err != nil {
				return uuuid.UUID{}
			}
			return uuid
		case []byte:
			uuid, err := uuuid.FromBytes(t)
			if err != nil {
				return uuuid.UUID{}
			}
			return uuid
		default:
			return uuuid.UUID{}
		}
	},
})
