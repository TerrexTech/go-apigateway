package resolver

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

type Asd struct {
	A int
}

type KaRespData struct {
	Inventory *[]model.Inventory `json:"inventory,omitempty"`
	Metric    *[]model.Metric    `json:"metric,omitempty"`
}

// EthyleneResolver is the resolver for Ethylene GraphQL query.
var SearchReport = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := "event.rns_eventstore.events"
	consTopic := "report.metric.reportmetric"

	searchReportJSON, err := json.Marshal(params.Args)
	log.Println(params.Args)
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	// if true {
	// 	// marg := params.Args["metric"].([]model.SearchParam)
	// 	iarg := params.Args["inventory"].([]model.SearchParam)
	// 	log.Printf("%+v", iarg, "&&&&&&&&&&&&&")
	// 	return string(searchReportJSON), nil
	// }

	log.Println(string(searchReportJSON))
	var _ = prodTopic

	rootValue := params.Info.RootValue.(map[string]interface{})
	ka := rootValue["kafkaFactory"].(*util.KafkaFactory)

	epio, err := ka.EnsureEventProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for RegisterResolver")
		return nil, err
	}
	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for cid")
		return nil, err
	}
	eventID, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error generating UUID for Register-Event")
		return nil, err
	}

	log.Println("===========CID:")
	log.Println(cid.String())

	// Publish Auth-Request on Kafka Topic
	go func() {
		epio.Input() <- &esmodel.Event{
			Action:        "query",
			CorrelationID: cid,
			AggregateID:   3,
			Data:          searchReportJSON,
			Timestamp:     time.Now(),
			UUID:          eventID,
			YearBucket:    2018,
			Version: 1,
		}
	}()
	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	krChan, err := ka.EnsureConsumerIO(consTopic, consTopic, false, cid)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for RegisterResolver")
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Timed out")
	case kr := <-krChan:
		log.Println("**********")
		queryInvMetResponse := handleInvMetResponse(kr, cid)
		log.Println(queryInvMetResponse, "^^^^^^^^^^^")
		if queryInvMetResponse != nil {
			return queryInvMetResponse, nil
		}
	}
	return nil, errors.New("Unknown Error")
	// return Asd{
	// 	A: 4,
	// }, nil
}

func handleInvMetResponse(
	kr esmodel.KafkaResponse,
	cid uuuid.UUID,
) *Out {
	log.Println("1111111111111111111111111111111")
	log.Println(string(kr.Result))
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "handleInvMetResponse: Error in KafkaResponse")
		// krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		// return &KaRespData{
		// 	Inventory: nil,
		// 	Metric:    nil,
		// }
	}

	log.Println("22222222222222222222222222222")
	searchResponse := KaRespData{}
	err := json.Unmarshal(kr.Result, &searchResponse)
	if err != nil {
		err = errors.Wrap(
			err,
			"handleInvMetResponse: Error while Unmarshalling report into KafkaResponse",
		)
		// 	log.Println(err)
		// krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		// return &KaRespData{
		// 	Inventory: nil,
		// 	Metric:    nil,
		// }
	}

	// log.Println("33333333333333333333333333333")
	log.Println(string(kr.Result), "&&&&&&&&&&")

	i, err := json.Marshal(searchResponse.Inventory)
	m, err := json.Marshal(searchResponse.Metric)

	log.Println(string(i))
	log.Println(string(m))

	return &Out{
		Inventory: string(i),
		Metric:    string(m),
	}
}

type Out struct {
	Inventory string
	Metric    string
}
