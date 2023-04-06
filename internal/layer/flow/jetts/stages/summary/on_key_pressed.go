package summary

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.STAR:
		c.ChangeStage <- "order"

	case hardware.SHARP:
		err := s.createTransaction(c)
		if err != nil {
			c.ChangeStage <- "emergency"
			return s.error(c, flow.ErrOutOfService, "out of service")
		}

		c.ChangeStage <- "payment_channel"
	}

	return nil
}
