package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func (s *Server) rpc(channel *Channel, msg amqp091.Delivery) error {
	if msg.CorrelationId == "" {
		return fmt.Errorf("rpc: no correlation ID")
	}

	routingKey := msg.Headers["routing-key"]
	if routingKey == nil {
		return fmt.Errorf("rpc: no routing Key")
	}

	key, ok := routingKey.(string)
	if !ok {
		return fmt.Errorf("rpc: routing key is not string: %t", routingKey)
	}

	handler := s.stacks[key]
	if handler == nil {
		return fmt.Errorf("rpc: no routing key registered")
	}

	ctx := NewContext(msg)
	err := (*handler)(ctx)
	if err != nil {
		return err
	}

	err = channel.PublishWithContext(
		ctx.UserContext,
		"",
		msg.ReplyTo,
		false,
		false,
		ctx.Publishing,
	)
	if err != nil {
		return err
	}

	return nil
}
