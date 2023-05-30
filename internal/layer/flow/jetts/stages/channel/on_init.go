package channel

import (
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	channels, err := s.paymentChannelRepo.FindMany(c.UserCtx, db.NewQuery().AddWhere("is_enable = ?", true))
	if err != nil {
		log.Error().Err(err).Msg("unable to get channel")
		c.ChangeStage <- flow.ORDER_STAGE
		return
	}
	if len(channels) == 0 {
		log.Error().Msg("no channel found")
		c.ChangeStage <- flow.ORDER_STAGE
		return
	}

	s.channels = channels
	s.frontendWs.SendPaymentChannel(c.UserCtx, channels)
}
