package payment

import (
	"fmt"
	"strings"

	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.bypass = false
	if !c.PaymentChannel.IsEnable {
		log.Error().Str("channel", c.PaymentChannel.Channel).Msg("unable to proceed payment, channel is disabled")
		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, fmt.Sprintf("%s is out of service", c.PaymentChannel.Channel))
		return
	}

	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "testpay":
		s.testpay(c)
	case "promptpay":
		s.promptpay(c)
	case "creditcard":
		go s.creditcard(c)
	default:
		return
	}
}
