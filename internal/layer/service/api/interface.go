package api

import (
	"context"

	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/domain/ksher"
	"vm-controller/internal/core/domain/link2500"
	"vm-controller/internal/layer/usecase/machine/model"
)

type Ksher interface {
	CreateOrder(context.Context, *entity.PaymentChannel, *ksher.CreateOrderBody) (*ksher.CreateOrderResult, error)
	CheckOrder(context.Context, *entity.PaymentChannel, string, *ksher.CheckOrderQuery) (*ksher.CheckOrderResult, error)
	RefundOrder(context.Context, *entity.PaymentChannel, string, *ksher.RefundOrderBody) (*ksher.RefundOrderResult, error)
}

type Link2500 interface {
	Sale(context.Context, *entity.PaymentChannel, *link2500.SaleRequest) (*link2500.SaleResult, error)
	Void(context.Context, *entity.PaymentChannel, *link2500.VoidRequest) (*link2500.VoidResult, error)
	Refund(context.Context, *entity.PaymentChannel, *link2500.RefundRequest) (*link2500.RefundResult, error)
	Settlement(context.Context, *entity.PaymentChannel, *link2500.SettlementRequest) (*link2500.SettlementResult, error)
}

type Topic interface {
	RegisterMachine(context.Context, *entity.Machine, *model.Machine) error
}
