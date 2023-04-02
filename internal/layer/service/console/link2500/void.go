package link2500_console

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
)

func (c *consoleImpl) Void(ctx context.Context, req *link2500.VoidRequest) (*link2500.VoidResult, error) {
	return nil, nil
}
