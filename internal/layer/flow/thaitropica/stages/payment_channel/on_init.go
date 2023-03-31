package payment_channel

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *Stage) OnInit(c *flow.Ctx) {
	channels, err := s.paymentChannelRepo.FindMany(c.UserCtx, []string{"active:=:true"})
	if err != nil {
		log.Error().Err(err).Msg("unable to get channel")
		return
	}

	s.channels = channels
	log.Info().Interface("channels", channels)
	s.ui.SendPaymentChannel(c.UserCtx, channels)
}
