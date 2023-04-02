package ksher_console

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (c *consoleImpl) CheckOrder(ctx context.Context, merchantOrderID string, query *ksher.CheckOrderQuery) (*ksher.CheckOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	result := ksher.CheckOrderResult{
		Note: "MOCKUP form console",
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
