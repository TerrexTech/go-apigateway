package resolver

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/TerrexTech/go-apigateway/gwerrors"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-apigateway/model"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

// authResponse is the GraphQL response on successful authentication.
type authResponse struct {
	authResponse *model.AuthResponse
	authErr      *gwerrors.KRError
}

// AccessTokenResolver is the resolver for AccessToken type.
var authHandler = func(
	ts auth.TokenStoreI,
	credentials map[string]interface{},
	pio *kafka.ProducerIO,
	cio *kafka.ConsumerIO,
) (interface{}, error) {
	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error generating UUID for cid")
		return nil, err
	}
	// Publish Auth-Request on Kafka Topic
	go func() {
		pio.ProducerInput() <- &esmodel.KafkaResponse{
			CorrelationID: cid,
			Input:         credentialsJSON,
			Topic:         pio.ID(),
		}
	}()

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

authResponseLoop:
	// Check auth-response messages for matching CorrelationID and return result
	for {
		select {
		case <-ctx.Done():
			break authResponseLoop
		case msg := <-cio.ConsumerMessages():
			cio.MarkOffset() <- msg
			authRes := handleAuthResponse(msg, ts, cid)
			if authRes != nil {
				if authRes.authErr == nil {
					return authRes.authResponse, nil
				}
				log.Println("ppppppppppppppppppppppp")
				return nil, errors.New(authRes.authErr.Error())
			}
		}
	}

	return nil, errors.New("Timed out")
}

func handleAuthResponse(
	msg *sarama.ConsumerMessage,
	ts auth.TokenStoreI,
	cid uuuid.UUID,
) *authResponse {
	user, krerr := parseKafkaResponse(msg, cid)
	if krerr != nil {
		err := errors.Wrap(krerr, "Error authenticating user")
		krerr.Err = err
		return &authResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}

	if user == nil {
		return nil
	}

	at, err := genAccessToken(user)
	if err != nil {
		err = errors.Wrap(err, "Error generating AccessToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "Error generating RefreshToken")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &authResponse{
			authResponse: nil,
			authErr:      krerr,
		}
	}

	return &authResponse{
		authResponse: &model.AuthResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
		authErr: nil,
	}
}

func parseKafkaResponse(
	msg *sarama.ConsumerMessage, cid uuuid.UUID,
) (*model.User, *gwerrors.KRError) {
	log.Println("aaaaaaaaaaaaaaaaaaaaaa")
	kr := &esmodel.KafkaResponse{}
	err := json.Unmarshal(msg.Value, kr)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error unmarshalling message into KafkaResponse")
		log.Println(err)
		return nil, nil
	}
	log.Println("zzzzzzzzzzzzzzzzzzzz")
	log.Printf("%+v", kr)
	if kr.AggregateID != 1 {
		return nil, nil
	}

	log.Println("bbbbbbbbbbbbbbbbbbbbbbb")
	if cid.String() != kr.CorrelationID.String() {
		log.Printf(
			"Error: Correlation ID Mistmatch: Expected CorrelationID: %s, Got: %s",
			cid.String(),
			kr.CorrelationID.String(),
		)
		return nil, nil
	}

	if kr.Error != "" {
		err = errors.Wrap(errors.New(kr.Error), "AuthKafkaResponseHandler Error")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return nil, krerr
	}

	log.Println("cccccccccccccccccccccccc")
	user := &model.User{}
	err = json.Unmarshal([]byte(kr.Result), user)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error unmarshalling message-result into User")
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return nil, krerr
	}

	log.Println("dddddddddddddddddddddddd")
	return user, nil
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
