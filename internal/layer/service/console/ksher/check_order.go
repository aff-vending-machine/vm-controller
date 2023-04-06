package ksher

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *consoleImpl) CheckOrder(ctx context.Context, channel *entity.PaymentChannel, merchantOrderID string, query *ksher.CheckOrderQuery) (*ksher.CheckOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := ksher.CheckOrderResult{
		Channel:         channel.Channel,
		Note:            "MOCKUP form console",
		Reference:       "MOCKUP",
		MerchantOrderID: merchantOrderID,
		Timestamp:       query.Timestamp,
		GatewayOrderID:  "p410760000",
		AcquirerOrderID: "90020230406180224936930",
	}
	if c.retry > 0 {
		result.ErrorCode = ksher.FAIL
		c.retry--
	} else {
		result.ErrorCode = ksher.SUCCESS
		c.retry = 5
	}

	return &result, nil
}
