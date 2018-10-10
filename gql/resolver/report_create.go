package resolver

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/TerrexTech/go-apigateway/gwerrors"
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-apigateway/util"
	esmodel "github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/uuuid"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

// EthyleneResolver is the resolver for Ethylene GraphQL query.
var CreateReportData = func(params graphql.ResolveParams) (interface{}, error) {
	prodTopic := os.Getenv("KAFKA_PRODUCER_TOPIC_REPORT")
	consGroup := os.Getenv("KAFKA_CONSUMER_GROUP_REPORT")
	consTopic := os.Getenv("KAFKA_CONSUMER_TOPIC_REPORT")

	rootValue := params.Info.RootValue.(map[string]interface{})
	kf := rootValue["kafkaFactory"].(*util.KafkaFactory)

	// Marshal Report
	reportJSON, err := json.Marshal(params.Args)
	if err != nil {
		err = errors.Wrap(err, "EthyleneResolver: Error marshalling ethylene report into JSON")
		return nil, err
	}

	// CorrelationID
	cid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "CreateDataResolver: Error generating UUID for cid")
		return nil, err
	}
	krpio, err := kf.EnsureKafkaResponseProducerIO(prodTopic, false)
	if err != nil {
		err = errors.Wrap(err, "Error creating ProducerIO for EthyleneResolver")
		return nil, err
	}
	// Publish Eth-Request on Kafka Topic
	go func() {
		krpio.Input() <- &esmodel.KafkaResponse{
			CorrelationID: cid,
			Input:         reportJSON,
			Topic:         prodTopic,
			AggregateID:   3,
		}
	}()

	cio, err := kf.EnsureConsumerIO(consGroup, consTopic, false, cid)
	if err != nil {
		err = errors.Wrap(err, "Error creating ConsumerIO for EthyleneResolver")
		return nil, err
	}
	// Timeout Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

reportResponseLoop:
	// Check ethylene-response messages for matching CorrelationID and return result
	for {
		select {
		case <-ctx.Done():
			break reportResponseLoop
		case msg := <-cio:
			reportResponse := handleGenDataReportResponse(msg, cid)
			if reportResponse != nil {
				if reportResponse.err == nil {
					return reportResponse.report, nil
				}
				return nil, errors.New(reportResponse.err.Error())
			}
			return &Resp1{
				Inventory: reportResponse.inventory,
				Metric:    reportResponse.metric,
				Report:    reportResponse.report,
			}, nil
		}
	}

	return nil, errors.New("Timed out")
}

type Resp1 struct {
	Inventory string
	Metric    string
	Report    string
}

type ReportResponse struct {
	report    string
	metric    string
	inventory string
	err       *gwerrors.KRError
}

type DataGen struct {
	Rdata []model.Report
	Mdata []model.Metric
	Idata []model.Inventory
}

func handleGenDataReportResponse(
	kr esmodel.KafkaResponse,
	cid uuuid.UUID,
) *ReportResponse {
	if kr.Error != "" {
		err := errors.New(kr.Error)
		err = errors.Wrap(err, "ReportResponseHandler: Error in KafkaResponse")
		krerr := gwerrors.NewKRError(err, kr.ErrorCode, err.Error())
		return &ReportResponse{
			report: "",
			err:    krerr,
		}
	}

	genData := &DataGen{}
	err := json.Unmarshal([]byte(kr.Result), genData)
	if err != nil {
		err = errors.Wrap(
			err,
			"ReportResponseHandler: Error while Unmarshalling report into KafkaResponse",
		)
		log.Println(err)
		krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
		return &ReportResponse{
			report: "",
			err:    krerr,
		}
	}

	var reportByte []byte
	var metricByte []byte
	var inventoryByte []byte

	for _, v := range genData.Rdata {
		reportByte, err = json.Marshal(&v)
		if err != nil {
			err = errors.Wrap(
				err,
				"ReportResponseHandler: Error while Unmarshalling report",
			)
			log.Println(err)
			krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
			return &ReportResponse{
				report: "",
				err:    krerr,
			}
		}
	}

	for _, v := range genData.Mdata {
		metricByte, err = json.Marshal(&v)
		if err != nil {
			err = errors.Wrap(
				err,
				"ReportResponseHandler: Error while Unmarshalling report",
			)
			log.Println(err)
			krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
			return &ReportResponse{
				report: "",
				err:    krerr,
			}
		}
	}

	for _, v := range genData.Idata {
		inventoryByte, err = json.Marshal(&v)
		if err != nil {
			err = errors.Wrap(
				err,
				"ReportResponseHandler: Error while Unmarshalling report",
			)
			log.Println(err)
			krerr := gwerrors.NewKRError(err, gwerrors.InternalError, err.Error())
			return &ReportResponse{
				report: "",
				err:    krerr,
			}
		}
	}

	reportJSON := string(reportByte)
	metricJSON := string(metricByte)
	inventoryJSON := string(inventoryByte)

	return &ReportResponse{
		report:    reportJSON,
		metric:    metricJSON,
		inventory: inventoryJSON,
		err:       nil,
	}
}
