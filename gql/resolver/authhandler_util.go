package resolver

import (
	"time"

	"github.com/TerrexTech/go-apigateway/gwerrors"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/pkg/errors"
)

// authResponse is the GraphQL response on successful authentication.
type authResponse struct {
	tokens *model.AuthTokens
	err    *gwerrors.KRError
}

func genAccessToken(user *model.User) (*model.AccessToken, error) {
	accessExp := 15 * time.Minute
	claims := &model.Claims{
		Role:      user.Role,
		Sub:       user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	accessToken, err := model.NewAccessToken(accessExp, claims)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error generating Access-Token")
		return nil, err
	}

	return accessToken, nil
}

func genRefreshToken(ts auth.TokenStoreI, user *model.User) (*model.RefreshToken, error) {
	refreshExp := (24 * 7) * time.Hour
	refreshToken, err := model.NewRefreshToken(refreshExp, user.UUID)
	if err != nil {
		err = errors.Wrap(err, "Error generating Refresh-Token")
		return nil, err
	}
	err = ts.Set(refreshToken)
	// We continue executing the code even if storing refresh-token fails since other parts
	// of application might still be accessible.
	if err != nil {
		err = errors.Wrapf(
			err,
			"Error storing RefreshToken in TokenStorage for UserID: %s", user.UUID,
		)
		return nil, err
	}

	return refreshToken, nil
}
