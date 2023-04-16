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

	s.bg(c)
	s.show(c)

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

	if c.Data.TotalQuantity() != c.Data.TotalReceived() {
		s.updateBrokenTransaction(c)
	} else {
		s.updateDoneTransaction(c)
	}

	s.status = DONE
	s.queue.Clear(c.UserCtx)

	c.ChangeStage <- "idle"
}
