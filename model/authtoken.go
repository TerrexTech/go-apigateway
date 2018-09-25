package model

//AuthResponse is the reponse for a successful authentication request.
type AuthResponse struct {
	AccessToken  *AccessToken  `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token"`
}
