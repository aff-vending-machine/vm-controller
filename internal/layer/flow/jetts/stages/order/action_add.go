package order

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"
)

func (s *stageImpl) actionAddItem(c *flow.Ctx, data item) error {
	if data.Quantity == 0 {
		return flow.ErrInvalidKey
	}

	var slot *entity.Slot
	for _, v := range s.slots {
		if v.Code == data.SlotCode {
			slot = &v
			break
		}
	}

	if slot == nil {
		return flow.ErrInvalidSlot
	}

	index := -1
	reserved := 0
	for i, item := range c.Data.Cart {
		if item.SlotCode == data.SlotCode {
			index = i
			reserved += item.Quantity
			break
		}
	}

	if slot.Stock < data.Quantity+reserved {
		return flow.ErrItemIsNotEnough
	}

	if index >= 0 {
		c.Data.Cart[index].Quantity = data.Quantity + reserved
	} else {
		c.Data.Cart = append(c.Data.Cart, hardware.Item{
			SlotCode: data.SlotCode,
			Name:     slot.Product.Name,
			ImageURL: slot.Product.ImageURL,
			Price:    slot.Product.Price,
			Quantity: data.Quantity,
			Received: 0,
		})
	}

	return nil
}
