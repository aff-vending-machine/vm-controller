package payment

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch key.ToString() {
	case hardware.StarKey:
		c.PaymentChannel = nil
		if s.ticker != nil {
			s.ticker.Stop()
		}
		c.ChangeStage <- "payment_channel"
	}

	return nil
}
