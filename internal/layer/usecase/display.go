package usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

type Display interface {
	Error(context.Context, error)
	Background(context.Context, string)
	Clear(context.Context)
	StageOrder(context.Context, hardware.Item, *hardware.Data)
	StageSummary(context.Context, []hardware.Item)
	StagePaymentChannel(context.Context, []entity.PaymentChannel)
	StagePaymentPromptPay(context.Context, string, float64)
	StagePaymentCreditCard(context.Context, float64)
	StageReceive(context.Context, []hardware.Item)
}
