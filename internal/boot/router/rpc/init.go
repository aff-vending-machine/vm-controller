package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
)

type server struct {
	*rabbitmq.Server
}

func New(app *rabbitmq.Wrapper) *server {
	return &server{
		rabbitmq.NewServer(app.Connection),
	}
}
