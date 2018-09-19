package resolver

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-apigateway/model"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/gofrs/uuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

type AuthResponse struct {
	ar  *model.AuthResponse
	err error
}

var LoginResolver = func(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	kio := rootValue["kafkaIO"].(*kafka.IO)
	ts := rootValue["tokenStore"].(auth.TokenStoreI)

	cid, err := uuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error generating UUID for cid")
		return nil, err
	}
	cidStr := cid.String()
	go func() {
		kio.ProducerInput() <- &esmodel.KafkaResponse{
			CorrelationID: cidStr,
			Input:         string(credentialsJSON),
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var arr *AuthResponse

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case msg := <-kio.ConsumerMessages():
			kr := &esmodel.KafkaResponse{}
			err := json.Unmarshal(msg.Value, kr)
			if err != nil {
				err = errors.Wrap(err, "Error unmarshalling message into KafkaResponse")
				kio.MarkOffset() <- msg
				log.Println(err)
				arr = &AuthResponse{
					ar:  nil,
					err: err,
				}
			}

			kio.MarkOffset() <- msg
			if cidStr != kr.CorrelationID {
				continue
			}
			kio.MarkOffset() <- msg

			user := &model.User{}
			err = json.Unmarshal([]byte(kr.Result), user)
			if err != nil {
				log.Println(err)
				kio.MarkOffset() <- msg
				arr = &AuthResponse{
					ar:  nil,
					err: err,
				}
				break
			}

			accessExp := 15 * time.Minute
			claims := &model.Claims{
				Role:      user.Role,
				Sub:       user.UUID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}
			accessToken, err := model.NewAccessToken(accessExp, claims)
			if err != nil {
				err = errors.Wrap(err, "Login Error: Error generating Access-Token")
				log.Println(err)
				kio.MarkOffset() <- msg
				arr = &AuthResponse{
					ar:  nil,
					err: err,
				}
				break
			}

			//=================RT
			refreshExp := (24 * 7) * time.Hour
			refreshToken, err := model.NewRefreshToken(refreshExp, user.UUID)
			if err != nil {
				err = errors.Wrap(err, "Error generating Refresh-Token")
				log.Println(err)
				kio.MarkOffset() <- msg
				arr = &AuthResponse{
					ar:  nil,
					err: err,
				}
				break
			}
			err = ts.Set(refreshToken)
			// We continue executing the code even if storing refresh-token fails since other parts
			// of application might still be accessible.
			if err != nil {
				err = errors.Wrapf(
					err,
					"Error storing RefreshToken in TokenStorage for UserID: %s", user.UUID,
				)
				log.Println(err)
				break
			}
			kio.MarkOffset() <- msg

			arr = &AuthResponse{
				ar: &model.AuthResponse{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
				},
				err: err,
			}
			break loop
		}
	}

	log.Println(arr)
	// select {
	// case <-ctx.Done():
	// 	return nil, errors.New("Login Error: Timed out")
	// case msg := <-kio.ConsumerMessages():

	// 	// }
	// 	// }(ctx)
	// }
	log.Println("==============")
	// producer.New
	// return auth.Login(db, redis, user)
	return arr.ar, nil
}
