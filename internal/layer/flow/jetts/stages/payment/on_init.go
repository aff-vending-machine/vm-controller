package payment

import (
	"strings"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.bg(c)
	s.show(c)

	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "promptpay":
		s.promptpay(c)
	case "creditcard":
		go s.creditcard(c)
	default:
		return
	}
}
