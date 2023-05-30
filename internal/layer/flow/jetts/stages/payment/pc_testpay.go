package payment

import (
	"time"

	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) testpay(c *flow.Ctx) {
	err := s.updateReferenceTransaction(c, "TEST01", "TEST02", "TEST03", "TEST04")
	if err != nil {
		c.ChangeStage <- flow.EMERGENCY_STAGE
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
					c.ChangeStage <- flow.EMERGENCY_STAGE
					return
				}

				s.frontendWs.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
				c.ChangeStage <- flow.RECEIVE_STAGE
				return
			}
			log.Debug().Int("count", count).Msg("polling check")
			count++

		case <-time.After(3 * time.Minute):
			s.frontendWs.SendError(c.UserCtx, "payment", "timeout")

			err := s.updateCancelTransaction(c, "machine")
			if err != nil {
				c.ChangeStage <- flow.EMERGENCY_STAGE
				return
			}

			c.ChangeStage <- flow.CHANNEL_STAGE
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}
