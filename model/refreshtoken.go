package model

import (
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type RefreshToken struct {
	Exp time.Time `json:"exp"`
	// Iat is the time the token was generated. This can be used in case the required
	// expiration time is different than Exp.
	Iat   time.Time `json:"iat"`
	Token string    `json:"token"`
	Sub   uuid.UUID `json:"sub"`
}

func NewRefreshToken(exp time.Duration, uid uuid.UUID) (*RefreshToken, error) {
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

func (rt *RefreshToken) Valid() bool {
	valid := time.Now().Before(rt.Exp)
	return valid
}

func (rt *RefreshToken) String() string {
	return rt.Token
}
