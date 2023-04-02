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
		s.createTransaction(c)
		c.ChangeStage <- "payment_channel"
	}

	return nil
}
