package model

//AuthResponse is the token provided to a user on successful login
type AuthResponse struct {
	AccessToken  *AccessToken  `json:"access_token"`
	RefreshToken *RefreshToken `json:"refresh_token"`
}
