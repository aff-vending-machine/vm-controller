package rpc

import "github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"

type Machine interface {
	Get(ctx *rabbitmq.Ctx) error
}

type Slot interface {
	Get(ctx *rabbitmq.Ctx) error
	Set(ctx *rabbitmq.Ctx) error
}

type Transaction interface {
	Get(ctx *rabbitmq.Ctx) error
	Clear(ctx *rabbitmq.Ctx) error
}
