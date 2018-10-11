package scalar

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/TerrexTech/go-apigateway/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var SearchParams = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "SeachParams",
	Description: "Search Parameters",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {
		switch t := value.(type) {
		case []model.SearchParam:
			s, err := json.Marshal(t)
			if err != nil {
				err = errors.Wrap(err, "SearchParams: Error Marshalling struct")
				log.Println(err)
				return ""
			}
			return string(s)
		case *[]model.SearchParam:
			s, err := json.Marshal(t)
			if err != nil {
				err = errors.Wrap(err, "SearchParams: Error Marshalling struct")
				log.Println(err)
				return ""
			}
			return string(s)
		default:
			return ""
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch t := value.(type) {
		case string:
			sp := &[]model.SearchParam{}
			err := json.Unmarshal([]byte(t), sp)
			if err != nil {
				return model.SearchParam{}
			}
			return *sp
		default:
			return model.SearchParam{}
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		value := valueAST.GetValue()
		switch t := value.(type) {
		case string:
			sp := &[]model.SearchParam{}
			err := json.Unmarshal([]byte(t), sp)
			if err != nil {
				return model.SearchParam{}
			}
			return *sp
		default:
			return model.SearchParam{}
		}
	},
})
