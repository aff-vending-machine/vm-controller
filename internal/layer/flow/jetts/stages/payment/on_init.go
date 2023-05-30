package payment

import (
	"fmt"
	"strings"

	"vm-controller/internal/core/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	if !c.PaymentChannel.IsEnable {
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
