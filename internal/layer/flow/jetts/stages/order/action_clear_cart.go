package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) actionClearCart(c *flow.Ctx) error {
	c.Data.Cart = make([]hardware.Item, 0)

	return nil
}
