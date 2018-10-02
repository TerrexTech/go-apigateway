package model

//AuthTokens is the reponse for a successful authentication request.
type AuthTokens struct {
	AccessToken  *AccessToken  `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token"`
}
