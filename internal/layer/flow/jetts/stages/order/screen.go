package order

import (
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) console(c *flow.Ctx, action string) {
	log.Info().Str("stage", "order").Str("action", action).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Float64("Price", c.Data.TotalPrice()).Msg("SLOG: order action")
}

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "order")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().Str("stage", "order").Int("step", s.step).Interface("pending", s.pendingItem).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Float64("Price", c.Data.TotalPrice()).Msg("SLOG: order action")

	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StageOrder(c.UserCtx, s.pendingItem, c.Data)
}

func (s *stageImpl) error(c *flow.Ctx, err error) error {
	log.Info().Str("stage", "order").Err(err).Msg("SLOG: order error")

	s.displayUc.Error(c.UserCtx, err)

	go func() {
		time.Sleep(5 * time.Second)
		if c.Stage == "order" {
			s.displayUc.Clear(c.UserCtx)
			s.displayUc.StageOrder(c.UserCtx, s.pendingItem, c.Data)
		}
	}()

	return err
}
