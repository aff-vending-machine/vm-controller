package payment

import (
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) testpay(c *flow.Ctx) {
	err := s.updateReferenceTransaction(c, "TEST01", "TEST02", "TEST03")
	if err != nil {
		c.ChangeStage <- "emergency"
		return
	}

	go s.pollingTestpay(c)
}

func (s *stageImpl) pollingTestpay(c *flow.Ctx) {
	count := 0
	s.ticker = time.NewTicker(5 * time.Second)
	defer s.ticker.Stop()

	for {
		select {
		case <-s.ticker.C:
			if c.Stage != "payment" || c.PaymentChannel.Channel != "testpay" {
				log.Error().Msg("cancelled by user")
				return
			}

			if count == 5 {
				err := s.updateTestTransaction(c)
				if err != nil {
					c.ChangeStage <- "emergency"
					return
				}

				c.ChangeStage <- "receive"
				return
			}
			log.Debug().Int("count", count).Msg("polling check")
			count++

		case <-time.After(3 * time.Minute):
			err := s.updateCancelTransaction(c, "machine")
			if err != nil {
				c.ChangeStage <- "emergency"
				return
			}

			c.ChangeStage <- "payment_channel"
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}