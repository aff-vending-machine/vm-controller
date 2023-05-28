package rpc

import "vm-controller/internal/core/infra/network/rabbitmq"

type Machine interface {
	Get(ctx *rabbitmq.Ctx) error
}

type PaymentChannel interface {
	Get(ctx *rabbitmq.Ctx) error
	Set(ctx *rabbitmq.Ctx) error
}

type Product interface {
	Set(ctx *rabbitmq.Ctx) error
}

type Slot interface {
	Get(ctx *rabbitmq.Ctx) error
	Set(ctx *rabbitmq.Ctx) error
}

type Transaction interface {
	Get(ctx *rabbitmq.Ctx) error
	Clear(ctx *rabbitmq.Ctx) error
}
