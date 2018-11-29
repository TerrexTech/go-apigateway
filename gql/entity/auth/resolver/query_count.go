package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// QueryCount returns the latest N inventory items.
var QueryCount = func(params graphql.ResolveParams) (interface{}, error) {
	result, err := genericQuery("count", params)
	if err != nil {
		err = errors.Wrap(err, "Error in QueryCount")
	}
	return result, err
}
