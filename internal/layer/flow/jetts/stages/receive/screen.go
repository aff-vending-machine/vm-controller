package receive

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "receive")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().
		Str("stage", "receive").
		Str("order_id", c.Data.MerchantOrderID).
		Interface("Cart", c.Data.Cart).
		Int("quantity", c.Data.TotalQuantity()).
		Int("received", c.Data.TotalReceived()).
		Int("price", int(c.Data.TotalPrice())).
		Int("pay", int(c.Data.TotalPay())).
		Msg("SLOG: receive action")

	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StageReceive(c.UserCtx, c.Data.Cart)
}

func (s *stageImpl) error(c *flow.Ctx, err error, msg string) error {
	log.Info().Str("stage", "receive").Err(err).Msg("SLOG: receive error")

	s.displayUc.Error(c.UserCtx, fmt.Errorf(msg))

	go func() {
		time.Sleep(5 * time.Second)
		s.show(c)
	}()

	return err
}
