package link2500

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *consoleImpl) Refund(ctx context.Context, channel *entity.PaymentChannel, req *link2500.RefundRequest) (*link2500.RefundResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := link2500.RefundResult{}

	return &result, nil
}