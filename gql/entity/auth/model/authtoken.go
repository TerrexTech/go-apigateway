package model

//AuthTokens contains AccessToken and RefreshToken for JWT authentication.
type AuthTokens struct {
	AccessToken  *AccessToken  `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token"`
}
