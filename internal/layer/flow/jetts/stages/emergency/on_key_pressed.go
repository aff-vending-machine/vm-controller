package emergency

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.SHARP:
		s.reset = s.reset + 1
		if s.reset > 7 {
			c.Reset()
			c.ChangeStage <- "idle"
		}
	}

	return nil
}
