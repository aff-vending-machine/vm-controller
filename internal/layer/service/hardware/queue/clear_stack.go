package queue

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (h *hardwareImpl) ClearStack(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	h.Clear(ctx)
	h.stacks = make(map[string]*hardware.Event)
}
