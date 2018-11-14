package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// QueryItem returns the matching items as per the paramters provided.
var QueryItem = func(params graphql.ResolveParams) (interface{}, error) {
	result, err := genericQuery("", params)
	if err != nil {
		err = errors.Wrap(err, "Error in QueryItem")
	}
	return result, err
}
