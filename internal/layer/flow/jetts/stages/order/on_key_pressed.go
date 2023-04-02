package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch s.step {
	case 0:
		return s.onFirstSlotInputStep(c, key)

	case 1:
		return s.onSlotInputStep(c, key)

	case 2:
		return s.onQuantityInputStep(c, key)
	}

	return nil
}
