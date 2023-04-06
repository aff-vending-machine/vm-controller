package link2500

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *consoleImpl) Void(ctx context.Context, channel *entity.PaymentChannel, req *link2500.VoidRequest) (*link2500.VoidResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := link2500.VoidResult{}

	return &result, nil
}
