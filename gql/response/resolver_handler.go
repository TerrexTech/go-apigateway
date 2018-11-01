package response

import "github.com/TerrexTech/go-apigateway/gwerrors"

// ResolverResponse is the return-value expected from handlers
// of GraphQL-Resolvers.
// This is only for convenience and optional to use, usage is
// still recommended for consistency.
// See handler-functions from exiting resolvers for examples.
type ResolverResponse struct {
	Result interface{}
	Err    *gwerrors.KRError
}
