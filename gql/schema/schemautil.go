package schema

import "github.com/graphql-go/graphql"

// createFields aggregates schemas from fields specified in entities.
func createFields(fields ...map[string]*graphql.Field) graphql.Fields {
	schema := graphql.Fields{}

	for _, fieldMap := range fields {
		for k, v := range fieldMap {
			schema[k] = v
		}
	}

	return schema
}
