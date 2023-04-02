package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	channels, err := s.paymentChannelRepo.FindMany(c.UserCtx, []string{"active:=:true"})
	if err != nil {
		log.Error().Err(err).Msg("unable to get channel")
		s.error(c, err, "out of service")
		c.ChangeStage <- "summary"
		return
	}

	s.channels = channels

	s.bg(c)
	s.show(c, channels)
}
