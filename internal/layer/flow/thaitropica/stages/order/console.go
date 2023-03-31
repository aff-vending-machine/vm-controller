package order

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *Stage) console(c *flow.Ctx, action string) {
	log.Info().Str("stage", "order").Str("action", action).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Float64("Price", c.Data.TotalPrice()).Msg("SLOG: order action")
}
