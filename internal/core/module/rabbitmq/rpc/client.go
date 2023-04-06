package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
)

type Client struct {
	Conn *rabbitmq.Connection
}

func NewClient(conn *rabbitmq.Connection) *Client {
	return &Client{
		Conn: conn,
	}
}
