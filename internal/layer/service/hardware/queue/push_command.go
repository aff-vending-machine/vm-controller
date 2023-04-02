package queue

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (h *hardwareImpl) PushCommand(ctx context.Context, key string, data string) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	cmd := h.client.LPush(ctx, key, data)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
