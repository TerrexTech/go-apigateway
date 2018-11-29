package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

func FormatResolvers(
	gqlParams graphql.ResolveParams,
	resolverMap map[string]graphql.FieldResolveFn,
) (interface{}, error) {
	result := make(map[string]interface{})

	if &gqlParams == nil {
		return nil, errors.New("nil gqlParams provided")
	}
	for k, v := range gqlParams.Args {
		if resolverMap[k] == nil {
			// This should never happen in production,
			// so we return error to let it be captured early in development.
			return nil, fmt.Errorf(`unable to find resolver for "%s"`, k)
		}

		// Set resolver-specific Args
		params := gqlParams
		args, assertOK := v.(map[string]interface{})
		if !assertOK {
			return nil, fmt.Errorf(
				`unable to assert args for "%s"; args must be map[string]interface{} type`, k,
			)
		}
		params.Args = args
		// Call the resolver-function
		resResult, err := (resolverMap[k])(params)
		if err != nil {
			return nil, err
		}
		result[k] = resResult
	}
	return result, nil
}
