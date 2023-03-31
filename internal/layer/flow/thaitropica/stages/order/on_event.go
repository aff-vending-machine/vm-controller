package order

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

func (s *Stage) OnEvent(c *flow.Ctx, event *hardware.Event) error {
	switch event.Status {
	case "00":
		return nil

	default:
		return nil
	}
}
