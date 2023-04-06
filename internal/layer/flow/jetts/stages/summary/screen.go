package summary

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "summary")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().
		Str("stage", "summary").
		Str("order_id", c.Data.MerchantOrderID).
		Interface("Cart", c.Data.Cart).
		Int("quantity", c.Data.TotalQuantity()).
		Int("price", int(c.Data.TotalPrice())).
		Msg("SLOG: summary action")

	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StageSummary(c.UserCtx, c.Data.Cart)
}

func (s *stageImpl) error(c *flow.Ctx, err error, msg string) error {
	log.Info().Str("stage", "summary").Err(err).Msg("SLOG: summary error")

	s.displayUc.Error(c.UserCtx, fmt.Errorf(msg))

	go func() {
		time.Sleep(5 * time.Second)
		if c.Stage == "summary" {
			s.show(c)
		}
	}()

	return err
}
