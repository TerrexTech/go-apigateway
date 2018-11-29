package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// QueryTimestamp is the resolver for Inventory GraphQL-query based on time-constraints.
var QueryTimestamp = func(params graphql.ResolveParams) (interface{}, error) {
	result, err := genericQuery("timestamp", params)
	if err != nil {
		err = errors.Wrap(err, "Error in QueryTimestamp")
	}
	return result, err
}
