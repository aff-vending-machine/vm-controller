package rabbitmq

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *Server) Listen(queue string) error {
	if s.Conn.IsClosed() {
		return fmt.Errorf("connection closed")
	}

	channel, err := s.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	q, err := channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	messages, err := channel.Consume(
		q.Name, // queue name
		"",     // name
		false,  // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var errs error
	for msg := range messages {
		log.Debug().Str("correlation_id", msg.CorrelationId).Str("key", msg.RoutingKey).Str("reply-to", msg.ReplyTo).Msg("rabbitmq: received")
		if msg.ReplyTo == "" {
			errs = s.topic(channel, msg)
		} else {
			errs = s.rpc(channel, msg)
		}

		if errs != nil {
			msg.Nack(false, false)
			log.Error().Err(errs).Msg("unable to process message delivery")
		} else {
			msg.Ack(false)
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
