package resolver

import (
	"time"

	"github.com/TerrexTech/go-apigateway/auth"
	gmodel "github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/uuuid"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// RegisterResolver is the resolver for Register GraphQL query.
var RegisterResolver = func(params graphql.ResolveParams) (interface{}, error) {
	// prodTopic := os.Getenv("EVENT_PRODUCER_TOPIC")
	// consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_REGISTER")

	rootValue := params.Info.RootValue.(map[string]interface{})
	db := rootValue["db"].(auth.DBI)

	uid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for event")
		return nil, err
	}
	// Create user from GraphQL Ags
	user := &gmodel.User{
		FirstName: params.Args["firstName"].(string),
		LastName:  params.Args["lastName"].(string),
		Email:     params.Args["email"].(string),
		Username:  params.Args["username"].(string),
		Password:  params.Args["password"].(string),
		Role:      params.Args["role"].(string),
		UUID:      uid,
	}
	user, err = db.Register(user)

	// AccessToken claims
	claims := &gmodel.Claims{
		Role:      user.Role,
		Sub:       user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	at, err := gmodel.NewAccessToken(15*time.Minute, claims)

	if err != nil {
		return nil, err
	}

	rt, err := gmodel.NewRefreshToken(7*24*time.Hour, uid)
	ar := &gmodel.AuthResponse{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return ar, nil
}
