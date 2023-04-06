package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq/rpc"
)

type server struct {
	*rpc.Server
}

func New(app *rabbitmq.Wrapper) *server {
	return &server{
		rpc.NewServer(app.Connection),
	}
}
