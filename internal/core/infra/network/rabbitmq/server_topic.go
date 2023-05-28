package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func (s *Server) topic(channel *Channel, msg amqp091.Delivery) error {
	handler := s.stacks[msg.RoutingKey]
	if handler == nil {
		return fmt.Errorf("topic: no routing key registered")
	}

	ctx := NewContext(msg)
	err := (*handler)(ctx)
	if err != nil {
		return err
	}

	return nil
}
