package resolver

import "github.com/TerrexTech/go-apigateway/gql"

// AuthResult is the result from registration or login.
type AuthResult struct {
	AccessToken  string    `json:"accessToken,omitempty"`
	RefreshToken string    `json:"refreshToken,omitempty"`
	Error        gql.Error `json:"error,omitempty"`
}
