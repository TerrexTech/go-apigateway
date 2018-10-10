package resolver

import (
	"encoding/json"
	"log"
	"time"

	"github.com/TerrexTech/go-apigateway/gwerrors"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

type SearchReportResponses struct {
	inventory *model.Inventory
	metric    *model.Metric
	err       *gwerrors.KRError
}

type SearchInvMet struct {
	Inventory model.Inventory
	Metric    model.Metric
}

// EthyleneResolver is the resolver for Ethylene GraphQL query.
var SearchReport = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := "event.rns_eventstore.events"
	// consTopic := report.metric.reportmetric

	searchReportJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "RegisterResolver: Error marshalling credentials into JSON")
		return nil, err
	}

	log.Println(searchReportJSON)
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
		}
	}()
	time.Sleep(1 * time.Second)
	return Asd{
		A: 4,
	}, nil
}

type Asd struct {
	A int
}

// type ReportResponse struct {
// 	report    string
// 	metric    string
// 	inventory string
// 	err       *gwerrors.KRError
// }

// type DataGen struct {
// 	Rdata []model.Report
// 	Mdata []model.Metric
// 	Idata []model.Inventory
// }

func handleInvMetResponse(
	kr esmodel.KafkaResponse,
	cid uuuid.UUID,
) *SearchReportResponses {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "handleInvMetResponse: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &SearchReportResponses{
			inventory: nil,
			metric:    nil,
			err:       krerr,
		}
	}

	searchInvMet := &SearchInvMet{}
	err := json.Unmarshal([]byte(kr.Result), searchInvMet)
	if err != nil {
		err = errors.Wrap(
			err,
			"handleInvMetResponse: Error while Unmarshalling report into KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &SearchReportResponses{
			inventory: nil,
			metric:    nil,
			err:       krerr,
		}
	}

	return &SearchReportResponses{
		inventory: &searchInvMet.Inventory,
		metric:    &searchInvMet.Metric,
		err:       nil,
	}
}
