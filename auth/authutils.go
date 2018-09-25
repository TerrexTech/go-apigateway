package auth

import (
	"fmt"
	"time"

	"github.com/TerrexTech/go-apigateway/auth/key"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/uuuid"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// RefreshAccessToken generates new AccessToken from provided RefreshToken.
func RefreshAccessToken(
	ts TokenStoreI,
	rt *model.RefreshToken,
	user *model.User,
) (*model.AccessToken, error) {
	if user.UUID != rt.Sub {
		return nil, errors.New("Error renewing AccessToken: User UUIDs don't match")
	}

	uid := rt.Sub
	receivedToken := rt.Token

	storedToken, err := ts.Get(uid)
	if err != nil {
		err = errors.Wrapf(
			err,
			"Error getting RefreshToken from Redis for UID: %s",
			uid.String(),
		)
		return nil, err
	}

	if receivedToken != storedToken {
		return nil, errors.New("Error renewing AccessToken: Invalid RefreshToken")
	}

	claims := &model.Claims{
		Role: user.Role,
		Sub:  user.UUID,
	}
	return model.NewAccessToken(15*time.Minute, claims)
}

// ParseAccessToken decodes the provided AccessToken into AccessToken model.
func ParseAccessToken(token string) (*model.AccessToken, error) {
	claims := jwt.MapClaims{}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, fmt.Errorf(
					"Error verifying AccessToken: Unexpected signing method: %v",
					t.Header["alg"],
				)
			}

			pubKey, err := key.GetPublicKey()
			if err != nil {
				err = errors.Wrap(err, "Error getting PublicKey")
				return nil, err
			}
			return pubKey, nil
		},
	)

	if err != nil {
		err = errors.Wrap(err, "Error parsing AccessToken")
		return nil, err
	}

	jti, err := uuuid.FromString(claims["jti"].(string))
	if err != nil {
		err = errors.Wrap(err, "Error parsing UUID from AccessToken")
		return nil, err
	}

	uid, err := uuuid.FromString(claims["sub"].(string))
	if err != nil {
		err = errors.Wrap(err, "Error parsing UserID from AccessToken")
		return nil, err
	}

	tokenClaims := &model.Claims{
		Role: claims["role"].(string),
		Sub:  uid,
	}

	exp := int64(claims["exp"].(float64))
	iat := int64(claims["iat"].(float64))
	at := &model.AccessToken{
		Claims:   tokenClaims,
		Exp:      time.Unix(exp, 0),
		Iat:      time.Unix(iat, 0),
		Jti:      jti,
		RawToken: parsedToken,
		Token:    token,
	}
	return at, nil
}
