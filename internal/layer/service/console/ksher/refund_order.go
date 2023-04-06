package ksher

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (*consoleImpl) RefundOrder(ctx context.Context, channel *entity.PaymentChannel, orderID string, body *ksher.RefundOrderBody) (*ksher.RefundOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := ksher.RefundOrderResult{
		ErrorCode: ksher.SUCCESS,
		Channel:   channel.Channel,
	}

	return &result, nil
}
