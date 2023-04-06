package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	c.Reset()
	c.ChangeStage <- "order"

	return nil
}
