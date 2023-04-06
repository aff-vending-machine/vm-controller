package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
	"github.com/rs/zerolog/log"
)

type Client struct {
	Conn *rabbitmq.Connection
	URL  string
}

func NewClient(url string) *Client {
	conn, err := rabbitmq.Dial(url)
	if err != nil {
		log.Error().Str("url", url).Err(err).Msg("unable to dial RabbitMQ")
	}

	return &Client{
		Conn: conn,
		URL:  url,
	}
}
