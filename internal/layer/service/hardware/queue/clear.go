package queue

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (h *hardwareImpl) Clear(ctx context.Context) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	cmd := h.client.FlushAll(ctx)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
