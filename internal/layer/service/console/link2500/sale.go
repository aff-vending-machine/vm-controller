package link2500

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *consoleImpl) Sale(ctx context.Context, channel *entity.PaymentChannel, req *link2500.SaleRequest) (*link2500.SaleResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := link2500.SaleResult{}

	return &result, nil
}
