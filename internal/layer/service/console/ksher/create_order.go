package ksher

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (*consoleImpl) CreateOrder(ctx context.Context, channel *entity.PaymentChannel, data *ksher.CreateOrderBody) (*ksher.CreateOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := ksher.CreateOrderResult{
		ErrorCode:       ksher.SUCCESS,
		Note:            "MOCKUP form console",
		Channel:         channel.Channel,
		Reference:       "MOCKUP",
		GatewayOrderID:  "p410760000",
		AcquirerOrderID: "90020230406180224936930",
		Reserved1:       "MOCKUP QRCODE",
		Amount:          data.Amount,
		MerchantOrderID: data.MerchantOrderID,
		Currency:        "THB",
		Timestamp:       data.Timestamp,
	}

	return &result, nil
}
