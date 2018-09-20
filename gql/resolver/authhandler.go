package resolver

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-apigateway/model"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type authResponse struct {
	authResponse *model.AuthResponse
	authErr      error
}

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
	cid, err := uuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "LoginResolver: Error generating UUID for cid")
		return nil, err
	}
	// Publish Auth-Request on Kafka Topic
	go func() {
		pio.ProducerInput() <- &esmodel.KafkaResponse{
			CorrelationID: cid.String(),
			Input:         string(credentialsJSON),
			Topic:         pio.ID(),
		}
	}()

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	authResChan := make(chan *authResponse)

authResponseLoop:
	// Check auth-response messages for matching CorrelationID and return result
	for {
		select {
		case <-ctx.Done():
			break authResponseLoop
		case msg := <-cio.ConsumerMessages():
			cio.MarkOffset() <- msg
			go handleAuthResponse(msg, ts, cid, authResChan)
			authRes := <-authResChan
			if authRes.authErr == nil {
				return authRes.authResponse, nil
			}
		}
	}

	return nil, errors.New("Timed out")
}

func handleAuthResponse(
	msg *sarama.ConsumerMessage,
	ts auth.TokenStoreI,
	cid uuid.UUID,
	outChan chan<- *authResponse,
) {
	user, err := parseKafkaResponse(msg, cid)
	if err != nil {
		err = errors.Wrap(err, "Error authenticating user")
		outChan <- &authResponse{
			authResponse: nil,
			authErr:      err,
		}
		return
	}

	at, err := genAccessToken(user)
	if err != nil {
		err = errors.Wrap(err, "Error generating AccessToken")
		outChan <- &authResponse{
			authResponse: nil,
			authErr:      err,
		}
		return
	}

	rt, err := genRefreshToken(ts, user)
	if err != nil {
		err = errors.Wrap(err, "Error generating RefreshToken")
		outChan <- &authResponse{
			authResponse: nil,
			authErr:      err,
		}
		return
	}

	outChan <- &authResponse{
		authResponse: &model.AuthResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
		authErr: nil,
	}
}

func parseKafkaResponse(msg *sarama.ConsumerMessage, cid uuid.UUID) (*model.User, error) {
	kr := &esmodel.KafkaResponse{}
	err := json.Unmarshal(msg.Value, kr)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error unmarshalling message into KafkaResponse")
		return nil, err
	}

	if cid.String() != kr.CorrelationID {
		return nil, errors.New("LoginResponseHandler: CorrelationID mistmatch")
	}

	user := &model.User{}
	err = json.Unmarshal([]byte(kr.Result), user)
	if err != nil {
		err = errors.Wrap(err, "LoginResponseHandler: Error unmarshalling message-result into User")
		return nil, err
	}

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
