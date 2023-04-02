package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) onQuantityInputStep(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.NUMBER:
		return s.step3_NumberKey(c, key)

	case hardware.STAR:
		return s.step3_StarKey(c)
	}

	return nil
}

func (s *stageImpl) step3_NumberKey(c *flow.Ctx, key hardware.Key) error {
	code := s.pendingItem.SlotCode
	quantity := key.ToNumber()

	if quantity == 0 {
		return s.error(c, flow.ErrEmptyItem)
	}

	index := -1
	reserved := 0
	for i, item := range c.Data.Cart {
		if item.SlotCode == code {
			index = i
			reserved += item.Quantity
			break
		}
	}

	if s.slot.Stock < quantity+reserved {
		return s.error(c, flow.ErrItemIsNotEnough)
	}

	if index >= 0 {
		c.Data.Cart[index].Quantity = quantity + reserved
	} else {
		c.Data.Cart = append(c.Data.Cart, hardware.Item{
			Name:     s.pendingItem.Name,
			SlotCode: code,
			Quantity: quantity,
			Price:    s.pendingItem.Price,
			Received: 0,
		})
	}

	s.reset()
	s.show(c)

	return nil
}

func (s *stageImpl) step3_StarKey(c *flow.Ctx) error {
	s.pendingItem.Quantity = 0
	s.pendingItem.Price = 0
	s.pendingItem.Name = ""
	s.pendingItem.SlotCode = s.pendingItem.SlotCode[:len(s.pendingItem.SlotCode)-1]
	s.backStep()
	s.show(c)

	return nil
}
