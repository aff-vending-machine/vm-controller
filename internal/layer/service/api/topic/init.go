package topic

import (
	"vm-controller/internal/core/infra/network/rabbitmq"
)

type apiImpl struct {
	*rabbitmq.Client
}

func New(conn *rabbitmq.Connection) *apiImpl {
	return &apiImpl{
		rabbitmq.NewClient(conn),
	}
}
