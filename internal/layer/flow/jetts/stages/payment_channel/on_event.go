package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) OnEvent(c *flow.Ctx, event *hardware.Event) error {
	switch event.Status {
	case "00":
		return nil

	default:
		return nil
	}
}
