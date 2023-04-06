package rpc

import (
	"github.com/rs/zerolog/log"
)

// Message properties
// content_type: Used to describe the mime-type of the encoding.
// reply_to: Commonly used to name a callback queue.
// correlation_id: Useful to correlate RPC responses with requests.
func (s *Server) Listen(queue string) {
	if s.Conn.IsClosed() {
		return
	}

	channel, err := s.Conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("unable to create channel")
		return
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
		log.Error().Err(err).Msg("unable to declare queue")
	}

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Error().Err(err).Msg("unable to config QOS")
	}

	messages, err := channel.Consume(
		q.Name, // queue name
		"rpc",  // name
		false,  // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // args
	)
	if err != nil {
		log.Error().Err(err).Msg("unable to create Consume")
	}

	for msg := range messages {
		if msg.ReplyTo == "" {
			log.Error().Interface("message", msg).Err(err).Msg("rpc: no ReplyTo")
			msg.Nack(false, false)
			continue
		}

		if msg.CorrelationId == "" {
			log.Error().Interface("message", msg).Err(err).Msg("rpc: no CorrelationId")
			msg.Nack(false, false)
			continue
		}

		routingKey := msg.Headers["routing-key"]

		if routingKey == nil {
			log.Error().Interface("message", msg).Err(err).Msg("rpc: routing key is header")
			msg.Nack(false, false)
			continue
		}

		key, ok := routingKey.(string)
		if !ok {
			log.Error().Interface("message", msg).Str("routingKey", key).Err(err).Msg("rpc: routing key is not string")
			msg.Nack(false, false)
			continue
		}

		handler := s.stacks[key]
		if handler == nil {
			log.Error().Interface("message", msg).Str("routingKey", key).Err(err).Msg("rpc: no routing key registered")
			msg.Nack(false, false)
			continue
		}

		ctx := NewContext(msg)
		if err := (*handler)(ctx); err != nil {
			log.Error().Str("routingKey", key).Interface("message", msg).Err(err).Msg("rpc: handler error")
			msg.Nack(false, false)
			continue
		}

		channel.PublishWithContext(
			ctx.UserContext,
			"",          // exchange
			msg.ReplyTo, // key
			false,       // mandatory
			false,       // immediate
			ctx.Publishing,
		)

		msg.Ack(false)
	}
}
