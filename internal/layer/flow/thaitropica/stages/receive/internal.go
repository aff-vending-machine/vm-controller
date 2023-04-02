package receive

import (
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

func makeCodeFilter(code string) []string {
	return []string{
		fmt.Sprintf("code:=:%s", code),
	}
}

func makeMerchantOrderIDFilter(id string) []string {
	return []string{
		fmt.Sprintf("merchant_order_id:=:%s", id),
	}
}

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
