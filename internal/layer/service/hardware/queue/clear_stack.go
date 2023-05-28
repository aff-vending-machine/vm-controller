package queue

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

func (h *hardwareImpl) ClearStack(ctx context.Context) {
	h.Clear(ctx)
	h.stacks = make(map[string]*hardware.Event)
}
