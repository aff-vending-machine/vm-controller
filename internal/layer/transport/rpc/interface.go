package rpc

import "github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq/rpc"

type Machine interface {
	Get(ctx *rpc.Ctx) error
}

type Slot interface {
	Get(ctx *rpc.Ctx) error
	Set(ctx *rpc.Ctx) error
}

type Transaction interface {
	Get(ctx *rpc.Ctx) error
	Clear(ctx *rpc.Ctx) error
}
