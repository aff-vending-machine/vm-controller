package machine_usecase

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

func (uc *usecaseImpl) Emergency(ctx context.Context) error {
	return uc.queueHw.Push(ctx, "EVENT", *hardware.NewEventFromString("000000000Z1"))
}
