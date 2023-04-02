package ksher_console

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (*consoleImpl) RefundOrder(ctx context.Context, orderID string, body *ksher.RefundOrderBody) (*ksher.RefundOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	return nil, fmt.Errorf("unimplementation")
}
