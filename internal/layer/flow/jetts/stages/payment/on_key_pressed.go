package payment

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.STAR:
		c.PaymentChannel = nil
		if s.ticker != nil {
			s.ticker.Stop()
		}
		c.ChangeStage <- "payment_channel"

	case hardware.SHARP:
		s.updateTestTransaction(c)
		c.ChangeStage <- "receive"
	}

	return nil
}
