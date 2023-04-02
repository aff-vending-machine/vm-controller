package payment

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "payment")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().
		Str("stage", "payment").
		Str("order_id", c.Data.MerchantOrderID).
		Str("payment_channel", c.PaymentChannel.Channel).
		Interface("Cart", c.Data.Cart).
		Int("quantity", c.Data.TotalQuantity()).
		Int("price", int(c.Data.TotalPrice())).
		Msg("SLOG: payment action")

	s.displayUc.Clear(c.UserCtx)
}

func (s *stageImpl) showCreditCard(c *flow.Ctx) {
	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StagePaymentCreditCard(c.UserCtx, c.Data.TotalPrice())
}

func (s *stageImpl) showPromptPay(c *flow.Ctx, res *ksher.CreateOrderResult) {
	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StagePaymentPromptPay(c.UserCtx, res.Reference, c.Data.TotalPrice())
}

func (s *stageImpl) error(c *flow.Ctx, err error, msg string) error {
	log.Info().Str("stage", "payment").Err(err).Msg("SLOG: payment error")

	s.displayUc.Error(c.UserCtx, fmt.Errorf(msg))

	go func() {
		time.Sleep(5 * time.Second)
		s.show(c)
	}()

	return err
}
