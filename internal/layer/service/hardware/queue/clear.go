package queue

import (
	"context"
)

func (h *hardwareImpl) Clear(ctx context.Context) error {
	cmd := h.client.FlushAll(ctx)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
