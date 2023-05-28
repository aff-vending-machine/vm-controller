package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/pkg/utils"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (c *Client) EmitRPC(ctx context.Context, queueTarget string, routingKey string, body []byte) ([]byte, error) {
	if c.Conn.IsClosed() {
		return nil, fmt.Errorf("lost rabbitmq connection")
	}

	channel, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	defer channel.QueuePurge(q.Name, false)
	defer channel.QueueDelete(q.Name, false, false, false)

	messages, err := channel.Consume(
		q.Name,         // queue name
		"rpc-receiver", // name
		true,           // autoAck
		false,          // exclusive
		false,          // noLocal
		false,          // noWait
		nil,            // args
	)
	if err != nil {
		return nil, err
	}

	corrId := utils.GenerateUUIDv4()

	log.Debug().Str("correlation_id", corrId).Str("key", routingKey).Str("reply_to", q.Name).Msg("rpc: emit")

	err = channel.PublishWithContext(
		ctx,
		"",          // exchange
		queueTarget, // key
		false,       // mandatory
		false,       // immediate
		amqp091.Publishing{
			Headers: amqp091.Table{
				"routing-key": routingKey,
			},
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
		},
	)

	if err != nil {
		return nil, err
	}

	select {
	case msg := <-messages:
		log.Debug().Str("correlation_id", msg.CorrelationId).Str("key", msg.RoutingKey).Msg("rpc: received")
		if corrId == msg.CorrelationId {
			return msg.Body, nil
		} else {
			return nil, errors.Wrapf(err, "id is mismatched, expect: %s, actual %s", corrId, msg.CorrelationId)
		}

	case <-time.After(15 * time.Second):
		return nil, errors.Wrapf(err, "%s", corrId)
	}
}
