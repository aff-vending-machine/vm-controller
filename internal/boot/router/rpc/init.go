package rpc

import (
	"vm-controller/internal/core/infra/network/rabbitmq"
)

type server struct {
	*rabbitmq.Server
}

func New(app *rabbitmq.Wrapper) *server {
	return &server{
		rabbitmq.NewServer(app.Connection),
	}
}
