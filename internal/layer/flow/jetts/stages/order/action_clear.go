package order

import "vm-controller/internal/core/flow"

func (s *stageImpl) actionClearItem(c *flow.Ctx, data item) error {
	for index, item := range c.Data.Cart {
		if item.SlotCode == data.SlotCode {
			if index+1 != len(c.Data.Cart) {
				c.Data.Cart = append(c.Data.Cart[:index], c.Data.Cart[index+1:]...)
			} else {
				c.Data.Cart = c.Data.Cart[:index]
			}
			break
		}
	}

	return nil
}
