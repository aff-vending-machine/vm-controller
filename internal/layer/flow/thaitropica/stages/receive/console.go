package receive

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) console(c *flow.Ctx) {
	log.Info().Str("stage", "receive").Int("remaining", len(c.Events)).Str("order_id", c.Data.MerchantOrderID).Interface("events", c.Events).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Float64("Price", c.Data.TotalPrice()).Float64("Pay", c.Data.TotalPay()).Msg("SLOG: receive event")
}
