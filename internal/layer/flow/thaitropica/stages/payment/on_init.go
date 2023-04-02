package payment

import (
	"fmt"
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	if !c.PaymentChannel.Active {
		s.ui.SendError(c.UserCtx, "payment", fmt.Sprintf("%s is out of service", c.PaymentChannel.Channel))
	}

	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "promptpay":
		s.promptpay(c)
	case "creditcard":
		go s.creditcard(c)
	case "wechatpay":
		s.wechatpay(c)
	case "alipay":
		s.alipay(c)
	default:
		return
	}
}
