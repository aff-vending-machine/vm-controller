package order

import (
	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) actionSetCart(c *flow.Ctx, data []item) error {
	c.Data.Cart = make([]hardware.Item, 0)
	for _, v := range s.slots {
		for _, d := range data {
			if v.Code != d.SlotCode {
				continue
			}
			if d.Quantity > 0 {
				index := -1
				reserved := 0
				for i, item := range c.Data.Cart {
					if item.SlotCode == d.SlotCode {
						index = i
						reserved += item.Quantity
						break
					}
				}

				if v.Stock < d.Quantity+reserved {
					return flow.ErrItemIsNotEnough
				}

				if index >= 0 {
					c.Data.Cart[index].Quantity = d.Quantity + reserved
				} else {
					c.Data.Cart = append(c.Data.Cart, hardware.Item{
						SlotCode: d.SlotCode,
						SKU:      v.Product.SKU,
						Name:     v.Product.Name,
						ImageURL: v.Product.ImageURL,
						Price:    v.Product.Price,
						Quantity: d.Quantity,
						Received: 0,
					})
				}
			} else {
				log.Warn().Str("slot_code", v.Code).Int("quantity", d.Quantity).Msg("quantity is 0")
			}
		}
	}

	return nil
}
