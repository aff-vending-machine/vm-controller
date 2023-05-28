package channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	channels, err := s.paymentChannelRepo.FindMany(c.UserCtx, db.NewQuery().AddWhere("active = ?", true))
	if err != nil {
		log.Error().Err(err).Msg("unable to get channel")
		c.ChangeStage <- "order"
		return
	}

	s.channels = channels
	s.frontendWs.SendPaymentChannel(c.UserCtx, channels)
}
