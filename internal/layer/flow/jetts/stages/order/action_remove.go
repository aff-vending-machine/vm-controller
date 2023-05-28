package order

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/flow"
)

func (s *stageImpl) actionRemoveItem(c *flow.Ctx, data item) error {
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
	for i, item := range c.Data.Cart {
		if item.SlotCode == data.SlotCode {
			index = i
			break
		}
	}

	if index >= 0 {
		if data.Quantity >= c.Data.Cart[index].Quantity {
			if index+1 != len(c.Data.Cart) {
				c.Data.Cart = append(c.Data.Cart[:index], c.Data.Cart[index+1:]...)
			} else {
				c.Data.Cart = c.Data.Cart[:index]
			}
		} else {
			c.Data.Cart[index].Quantity = c.Data.Cart[index].Quantity - data.Quantity
		}
	} else {
		return flow.ErrInvalidSlot
	}

	return nil
}
