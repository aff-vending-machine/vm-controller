package order

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

func (s *Stage) actionClearCart(c *flow.Ctx) error {
	c.Data.Cart = make([]hardware.Item, 0)

	return nil
}
