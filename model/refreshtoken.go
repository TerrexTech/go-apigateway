package model

import (
	"math/rand"
	"time"

	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

// RefreshToken represents a JWT-auth RefreshToken.
type RefreshToken struct {
	Exp time.Time `json:"exp"`
	// Iat is the time the token was generated. This can be used in case the required
	// expiration time is different than Exp.
	Iat   time.Time `json:"iat"`
	Token string    `json:"token"`
	// Sub is the UserUUID for which the RefreshToken was generated.
	Sub uuuid.UUID `json:"sub"`
}

// NewRefreshToken generates a new RefreshToken.
func NewRefreshToken(exp time.Duration, uid uuuid.UUID) (*RefreshToken, error) {
	if uid.String() == "" {
		return nil, errors.New("Error generating new AccessToken: User uuid absent")
	}

	// Length of token
	const size = 64
	// Characters to use for generating token
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	const charLength = len(chars)

	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(charLength)]
	}

	token := string(bytes)
	return &RefreshToken{
		Exp:   time.Now().Add(exp),
		Iat:   time.Now(),
		Token: token,
		Sub:   uid,
	}, nil
}

// Valid returns true if the Refresh token is valid, and hasn't expired.
func (rt *RefreshToken) Valid() bool {
	valid := time.Now().Before(rt.Exp)
	return valid
}

// String returns the string representation of RefreshToken.
func (rt *RefreshToken) String() string {
	return rt.Token
}
