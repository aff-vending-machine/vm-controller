package machine

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
)

func (uc *usecaseImpl) OpenGate(ctx context.Context) error {
	return uc.queueHw.Push(ctx, "EVENT", *hardware.NewEventFromString("000000000Z2"))
}
