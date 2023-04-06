package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) onFirstSlotInputStep(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.NUMBER:
		s.step1_NumberKey(c, key)

	case hardware.ALPHABET:
		s.step1_AlphabetKey(c, key)

	case hardware.STAR:
		s.step1_StarKey(c)

	case hardware.SHARP:
		s.step1_SharpKey(c)
	}

	return nil
}

func (s *stageImpl) step1_NumberKey(c *flow.Ctx, key hardware.Key) {
	s.pendingItem.SlotCode = "0" + key.ToString()
	s.nextStep()
	s.show(c)
}

func (s *stageImpl) step1_AlphabetKey(c *flow.Ctx, key hardware.Key) {
	s.pendingItem.SlotCode = key.ToString()
	s.nextStep()
	s.show(c)
}

func (s *stageImpl) step1_StarKey(c *flow.Ctx) {
	size := len(c.Data.Cart)
	if size > 0 {
		c.Data.Cart = c.Data.Cart[:size-1]
		s.show(c)
	} else {
		s.reset()
		c.ChangeStage <- "idle"
	}
}

func (*stageImpl) step1_SharpKey(c *flow.Ctx) {
	if len(c.Data.Cart) > 0 {
		c.ChangeStage <- "summary"
	}
}
