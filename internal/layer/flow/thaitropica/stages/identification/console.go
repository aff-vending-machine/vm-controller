package identification

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) console_processing(c *flow.Ctx) {
	log.Info().Str("stage", "identification").Str("order_id", c.Data.MerchantOrderID).Msg("SLOG: processing")
}

func (s *stageImpl) console_email_used(c *flow.Ctx, email string) {
	log.Info().Str("stage", "identification").Str("order_id", c.Data.MerchantOrderID).Str("email", email).Msg("SLOG: email is used")
}

func (s *stageImpl) console_request(c *flow.Ctx) {
	log.Info().Str("stage", "identification").Str("order_id", c.Data.MerchantOrderID).Interface("OTP stack", s.stacks).Msg("SLOG: request OTP from user")
}

func (s *stageImpl) console_check(c *flow.Ctx, ref string, otp string, result bool) {
	log.Info().Str("stage", "identification").Str("order_id", c.Data.MerchantOrderID).Interface("OTP stack", s.stacks).Str("reference", ref).Str("OTP", otp).Bool("result", true).Msg("SLOG: check OTP from user")
}
