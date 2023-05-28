package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) actionSetItem(c *flow.Ctx, data item) error {
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

	if slot.Stock < data.Quantity {
		return flow.ErrItemIsNotEnough
	}

	for index, item := range c.Data.Cart {
		if item.SlotCode == data.SlotCode {
			c.Data.Cart[index].Quantity = data.Quantity
			return nil
		}
	}

	c.Data.Cart = append(c.Data.Cart, hardware.Item{
		SlotCode: data.SlotCode,
		Name:     slot.Product.Name,
		ImageURL: slot.Product.ImageURL,
		Price:    slot.Product.Price,
		Quantity: data.Quantity,
		Received: 0,
	})

	return nil
}
