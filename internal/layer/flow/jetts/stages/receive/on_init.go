package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.queue.Clear(c.UserCtx)

	if err := s.addEvents(c); err != nil {
		log.Error().Err(err).Msg("unable to add events")
		return
	}

	log.Info().Str("stage", "receive").Int("remaining", len(c.Events)).Str("order_id", c.Data.MerchantOrderID).Interface("events", c.Events).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Float64("Price", c.Data.TotalPrice()).Float64("Pay", c.Data.TotalPay()).Msg("SLOG: receive event")
	go s.checkEvent(c)
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

func (s *stageImpl) checkEvent(c *flow.Ctx) {
	total := c.Data.TotalQuantity()

	log.Info().Msg("start polling")
	s.polling = true
	s.queue.Polling(c.UserCtx, "RESPONSE", total, s.feedback(c))
	s.polling = false
	log.Info().Msg("stop polling")

	if s.status == CANCEL {
		s.updateCancelTransaction(c)

		s.queue.Clear(c.UserCtx)
		c.ChangeStage <- "idle"
		return
	}

	if c.Data.TotalQuantity() != c.Data.TotalReceived() {
		s.updateMachineFailedTransaction(c)
	} else {
		s.updateDoneTransaction(c)
	}

	log.Info().Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Msg("DONE")
	s.frontendWs.SendDone(c.UserCtx, c.Data.MerchantOrderID, c.Data.Cart)

	s.status = DONE
	s.queue.Clear(c.UserCtx)

	c.ChangeStage <- "idle"
}
