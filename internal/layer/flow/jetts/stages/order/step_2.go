package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) onSlotInputStep(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.NUMBER:
		return s.step2_NumberKey(c, key)

	case hardware.STAR:
		return s.step2_StarKey(c, key)
	}

	return nil
}

func (s *stageImpl) step2_NumberKey(c *flow.Ctx, key hardware.Key) error {
	s.pendingItem.SlotCode += key.ToString()
	if len(s.pendingItem.SlotCode) != 3 {
		return nil
	}

	filters := makeSlotFilter(s.pendingItem.SlotCode)
	slot, err := s.slotRepo.FindOne(c.UserCtx, filters)
	if err != nil {
		s.reset()
		return s.error(c, flow.ErrNoItem)
	}

	if slot.Stock == 0 {
		s.reset()
		return s.error(c, flow.ErrEmptyItem)
	}

	if !slot.IsEnable {
		s.reset()
		return s.error(c, flow.ErrInvalidSlot)
	}

	s.pendingItem.Name = slot.Product.Name
	s.pendingItem.Price = slot.Product.Price
	s.slot = slot

	s.nextStep()
	s.show(c)

	return nil
}

func (s *stageImpl) step2_StarKey(c *flow.Ctx, key hardware.Key) error {
	size := len(s.pendingItem.SlotCode)
	s.pendingItem.SlotCode = s.pendingItem.SlotCode[:size-1]
	size = len(s.pendingItem.SlotCode)
	if size == 0 || s.pendingItem.SlotCode == "0" {
		s.reset()
	}
	s.show(c)

	return nil
}
