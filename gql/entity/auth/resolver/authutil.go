package resolver

import (
	"time"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/entity/auth/model"
	"github.com/pkg/errors"
)

func genAccessToken(user *model.User) (*model.AccessToken, error) {
	accessExp := 15 * time.Minute
	claims := &model.Claims{
		Role:      user.Role,
		Sub:       user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	accessToken, err := model.NewAccessToken(accessExp, claims)
	if err != nil {
		err = errors.Wrap(err, "AuthGenAccessToken: Error generating Access-Token")
		return nil, err
	}

	return accessToken, nil
}

func genRefreshToken(ts auth.TokenStoreI, user *model.User) (*model.RefreshToken, error) {
	refreshExp := (24 * 7) * time.Hour
	refreshToken, err := model.NewRefreshToken(refreshExp, user.UserID)
	if err != nil {
		err = errors.Wrap(err, "AuthGenRefreshToken: Error generating Refresh-Token")
		return nil, err
	}
	err = ts.Set(refreshToken)
	if err != nil {
		err = errors.Wrapf(
			err,
			"AuthGenRefreshToken: Error storing RefreshToken in TokenStorage for UserID: %s", user.UserID,
		)
		return nil, err
	}

	return refreshToken, nil
}
