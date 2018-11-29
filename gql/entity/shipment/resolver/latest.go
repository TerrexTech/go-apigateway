package resolver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/gql"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-common-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

var Latest = func(params graphql.ResolveParams) (interface{}, error) {
	consTopic := os.Getenv("KC_TOPIC_SHIPMENT_ITEMS")

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// Marshal User-credentials
	credentialsJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error generating UUID for cid")
		return nil, err
	}
	uuid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error generating UUID for Register-Event")
		return nil, err
	}
	kf.CmdProducer() <- &esmodel.Command{
		Action:        "Latest",
		CorrelationID: cid,
		Data:          credentialsJSON,
		ResponseTopic: "gateway.response",
		SourceTopic:   consTopic,
		Source:        "go-apigateway",
		Timestamp:     time.Now().UTC().Unix(),
		TTLSec:        5,
		UUID:          uuid,
	}

	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	docChan, err := kf.EnsureConsumerIO("gateway.response", "gateway.response", false, uuid)
	if err != nil {
		err = errors.Wrap(err, "AuthRegisterResolver: Error creating ConsumerIO")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case doc := <-docChan:
		registerResp := handleLatestResponse(doc)
		return registerResp, nil
	}
}

func handleLatestResponse(doc esmodel.Document) map[string]interface{} {
	if doc.Error != "" {
		err := errors.New(doc.Error)
		err = errors.Wrap(err, "AuthLoginResponseHandler: Error in Document")
		log.Println(err)
		return map[string]interface{}{
			"error": gql.Error{
				Code:    doc.ErrorCode,
				Message: doc.Error,
			},
		}
	}

	items := []map[string]interface{}{}
	err := json.Unmarshal(doc.Data, &items)
	if err != nil {
		err = errors.Wrap(
			err,
			"Error while unmarshalling document into Item",
		)
		log.Println(err)
		return map[string]interface{}{
			"error": gql.Error{
				Code:    esmodel.InternalError,
				Message: err.Error(),
			},
		}
	}
	return map[string]interface{}{
		"items": items,
	}
}
