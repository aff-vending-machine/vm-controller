package ksher_console

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (*consoleImpl) CreateOrder(ctx context.Context, data *ksher.CreateOrderBody) (*ksher.CreateOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := ksher.CreateOrderResult{
		ErrorCode: ksher.SUCCESS,
		Note:      "MOCKUP form console",
		Reserved1: "MOCKUP QRCODE",
	}

	return &result, nil
}
