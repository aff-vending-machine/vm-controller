package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	channels, err := s.paymentChannelRepo.FindMany(c.UserCtx, []string{"active:=:true.(bool)"})
	if err != nil {
		log.Error().Err(err).Msg("unable to get channel")
		s.error(c, err, "out of service")
		c.ChangeStage <- "order"
		return
	}

	s.channels = channels
	s.frontendWs.SendPaymentChannel(c.UserCtx, channels)

	s.bg(c)
	s.show(c, channels)
}
