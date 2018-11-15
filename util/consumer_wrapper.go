package util

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/Shopify/sarama"
	"github.com/TerrexTech/go-kafkautils/kafka"
	"github.com/pkg/errors"
)

type responseConsumer struct {
	consumer    *kafka.Consumer
	topic       string
	consumerCtx context.Context
	errGroup    *errgroup.Group

	msgChan     chan *sarama.ConsumerMessage
	isConsuming bool
}

func newResponseConsumer(
	ctx context.Context,
	errGroup *errgroup.Group,
	config *kafka.ConsumerConfig,
) (*responseConsumer, error) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		err = errors.Wrap(err, "Error while creating ResponseConsumer")
		return nil, err
	}
	return &responseConsumer{
		consumer:    consumer,
		topic:       config.Topics[0],
		consumerCtx: ctx,
		errGroup:    errGroup,

		msgChan:     make(chan *sarama.ConsumerMessage),
		isConsuming: false,
	}, nil
}

func (*responseConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (*responseConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *responseConsumer) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for {
		select {
		case <-r.consumerCtx.Done():
			log.Printf("--> Closing Consumer on Topic: %s", r.topic)
			close(r.msgChan)
			return errors.New("service-context closed")
		case msg := <-claim.Messages():
			log.Println(string(msg.Value))
			r.msgChan <- msg
		}
	}
}

func (r *responseConsumer) Messages() <-chan *sarama.ConsumerMessage {
	if !r.isConsuming {
		r.errGroup.Go(func() error {
			r.isConsuming = true
			err := r.consumer.Consume(r.consumerCtx, r)
			if err != nil {
				err = errors.Wrap(
					err,
					"ResponseConsumer: Error while attempting to consume messages",
				)
				log.Println(err)
				return err
			}
			return nil
		})
	}
	return r.msgChan
}
