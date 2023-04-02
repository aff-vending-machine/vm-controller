package api

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
)

type Ksher interface {
	CreateOrder(context.Context, *entity.PaymentChannel, *ksher.CreateOrderBody) (*ksher.CreateOrderResult, error)
	CheckOrder(context.Context, *entity.PaymentChannel, string, *ksher.CheckOrderQuery) (*ksher.CheckOrderResult, error)
	RefundOrder(context.Context, *entity.PaymentChannel, string, *ksher.RefundOrderBody) (*ksher.RefundOrderResult, error)
}

type Link2500 interface {
	Sale(context.Context, *entity.PaymentChannel, *link2500.SaleRequest) (*link2500.SaleResult, error)
	Void(context.Context, *entity.PaymentChannel, *link2500.VoidRequest) (*link2500.VoidResult, error)
	Settlement(context.Context, *entity.PaymentChannel) (*link2500.SettlementResult, error)
}
