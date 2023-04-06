package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
)

type Server struct {
	Conn   *rabbitmq.Connection
	stacks map[string]*Handler
}

func NewServer(conn *rabbitmq.Connection) *Server {
	return &Server{
		Conn:   conn,
		stacks: make(map[string]*Handler),
	}
}
