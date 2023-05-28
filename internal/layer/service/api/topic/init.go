package topic

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
)

type apiImpl struct {
	*rabbitmq.Client
}

func New(conn *rabbitmq.Connection) *apiImpl {
	return &apiImpl{
		rabbitmq.NewClient(conn),
	}
}
