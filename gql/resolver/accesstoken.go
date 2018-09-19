package resolver

import (
	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/gofrs/uuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

var AccessTokenResolver = func(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	redis := rootValue["redis"].(auth.TokenStoreI)

	rtStr := params.Args["refreshToken"].(string)
	uid := params.Args["sub"].(string)

	parsedUID, err := uuid.FromString(uid)
	if err != nil {
		return nil, errors.New("Error parsing RefreshToken: Cannot parse Sub")
	}

	rt := &model.RefreshToken{
		Sub:   parsedUID,
		Token: rtStr,
	}

	//==============Fix below
	user, err := &model.User{}, nil
	if err != nil {
		return nil, errors.Wrap(
			err,
			"Error parsing RefreshToken: Cannot get user with specified UUID",
		)
	}
	at, err := auth.RefreshAccessToken(redis, rt, user)
	if err != nil {
		err = errors.Wrap(err, "Error renewing AccessToken")
		return nil, err
	}
	return &model.AuthResponse{
		AccessToken: at,
	}, nil
}
