package payment

import (
	"fmt"
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.bg(c)
	s.show(c)

	if !c.PaymentChannel.Active {
		s.frontendWs.SendError(c.UserCtx, "payment", fmt.Sprintf("%s is out of service", c.PaymentChannel.Channel))
	}

	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "promptpay":
		s.promptpay(c)
	case "creditcard":
		go s.creditcard(c)
	default:
		return
	}
}
