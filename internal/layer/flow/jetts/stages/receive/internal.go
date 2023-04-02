package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) addEvents(c *flow.Ctx) error {
	for _, item := range c.Data.Cart {
		for index := 0; index < item.Quantity; index++ {
			event := hardware.NewEvent(index, item)
			err := s.queue.Push(c.UserCtx, "QUEUE", event)
			if err != nil {
				return err
			}

			c.AddWaitingEvent(event)
		}
	}
	return nil
}
