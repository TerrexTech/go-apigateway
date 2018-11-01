package scalar

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/pkg/errors"
)

var Map = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Map",
	Description: "A Type representing a map, analogous to map[string]interface{}",
	// Serialize serializes `map` to json-string.
	Serialize: func(value interface{}) interface{} {
		switch value.(type) {
		case map[string]interface{}:
			jsonStr, err := json.Marshal(value)
			if err != nil {
				err = errors.Wrap(err, "Error while marshalling GraphQL Map-Scalar")
				log.Println(err)
				return nil
			}
			return string(jsonStr)
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `map`.
	ParseValue: func(value interface{}) interface{} {
		result := map[string]interface{}{}

		input := []byte{}
		switch value.(type) {
		case string:
			input = []byte(value.(string))
		case []byte:
			input = value.([]byte)
		default:
			return nil
		}

		err := json.Unmarshal(input, &result)
		if err != nil {
			err = errors.Wrap(err, "Error while unmarshalling GraphQL Map-Scalar")
			log.Println(err)
			return nil
		}
		return result
	},
	// ParseLiteral parses GraphQL AST value to `map`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		value := valueAST.GetValue()
		result := map[string]interface{}{}

		input := []byte{}
		switch value.(type) {
		case string:
			input = []byte(value.(string))
		case []byte:
			input = value.([]byte)
		default:
			return nil
		}

		err := json.Unmarshal(input, &result)
		if err != nil {
			err = errors.Wrap(err, "Error while unmarshalling GraphQL Map-Scalar")
			log.Println(err)
			return nil
		}
		return result
	},
})
