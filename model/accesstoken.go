package model

import (
	"time"

	"github.com/TerrexTech/go-apigateway/auth/key"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// AccessToken is used to authenticate with other services.
type AccessToken struct {
	// Claims represent the JWT Claims.
	Claims *Claims `json:"claims"`
	// Exp is the time the token is intended to expire.
	Exp time.Time `json:"exp"`
	// Iat is the time the token was generated. This can be used in case the required
	// expiration time is different than Exp.
	Iat time.Time `json:"iat"`
	// Jti is the unique-identifier for this token.
	Jti uuid.UUID `json:"jti"`
	// RawToken is the actual/un-encoded JWT RawToken.
	RawToken *jwt.Token `json:"raw_token"`
	// Token is the JWT token ecoded with required claims.
	// This exists for caching purposes.
	Token string `json:"token"`
}

type Claims struct {
	// Role represent's user's authorization-level.
	Role string `json:"role"`
	// Sub is the user-uuid for whom the token is generated.
	Sub       uuid.UUID `json:"sub"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

func NewAccessToken(exp time.Duration, claims *Claims) (*AccessToken, error) {
	if claims.Sub.String() == "" {
		return nil, errors.New("Error generating new AccessToken: Sub absent")
	}
	if claims.Role == "" {
		return nil, errors.New("Error generating new AccessToken: Role absent")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		err = errors.Wrap(
			err,
			"Error generating new AccessToken: Error generating UUID",
		)
		return nil, err
	}

	at := &AccessToken{
		Claims: claims,
		Exp:    time.Now().Add(exp),
		Iat:    time.Now(),
		Jti:    uuid,
	}

	encoded, err := at.Encode()
	if err != nil {
		err = errors.Wrap(
			err,
			"Error generating new AccessToken: Error encoding AccessToken",
		)
		return nil, err
	}
	at.Token = encoded

	return at, nil
}

func (at *AccessToken) Encode() (string, error) {
	var err error
	if at.Exp.Before(time.Now()) {
		return "", errors.New("Token encode-error: Exp time cannot be before current time")
	}

	if at.Token == "" {
		token := jwt.New(jwt.SigningMethodRS512)
		token.Valid = true
		token.Claims = jwt.MapClaims{
			"exp":        at.Exp.Unix(),
			"iat":        at.Iat.Unix(),
			"jti":        at.Jti.String(),
			"sub":        at.Claims.Sub.String(),
			"role":       at.Claims.Role,
			"first_name": at.Claims.FirstName,
			"last_name":  at.Claims.LastName,
		}
		at.RawToken = token

		privKey, err := key.GetPrivateKey()
		if err != nil {
			err = errors.Wrap(err, "Error getting PrivateKey for signing AccessToken")
			return "", err
		}

		at.Token, err = token.SignedString(privKey)
		if err != nil {
			err = errors.Wrap(err, "Token encode-error: Error signing token")
		}
	}
	return at.Token, err
}

func (at *AccessToken) Valid() bool {
	valid := time.Now().Before(at.Exp)
	if !valid {
		at.RawToken.Valid = false
	}
	return valid
}

func (at *AccessToken) String() string {
	return at.Token
}
