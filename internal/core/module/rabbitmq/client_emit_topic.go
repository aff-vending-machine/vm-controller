package rabbitmq

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/pkg/utils"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (c *Client) EmitTopic(ctx context.Context, exchange string, queue string, routingKey string, data []byte) error {
	if c.Conn.IsClosed() {
		return fmt.Errorf("lost rabbitmq connection")
	}

	channel, err := c.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	corrId := utils.GenerateUUIDv4()

	log.Debug().Str("correlation_id", corrId).Str("exchange", exchange).Str("Queue", queue).Str("routingKey", routingKey).Msg("topic: emit")

	err = channel.Bind(exchange, queue, routingKey)
	if err != nil {
		return err
	}
	defer channel.Unbind(exchange, queue, routingKey)

	err = channel.PublishWithContext(
		ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp091.Publishing{
			Headers: amqp091.Table{
				"routing-key": routingKey,
			},
			CorrelationId: corrId,
			Body:          data,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
