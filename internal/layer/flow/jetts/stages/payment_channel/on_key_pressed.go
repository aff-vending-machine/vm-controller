package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch key.ToString() {
	case hardware.NUMBER:
		index := key.ToNumber() - 1
		if index >= 0 && index < len(s.channels) {
			err := s.updateTransaction(c, s.channels[index])
			if err != nil {
				c.ChangeStage <- "emergency"
				return s.error(c, flow.ErrInvalidKey, "out of service")
			}

			c.PaymentChannel = &s.channels[index]
			c.ChangeStage <- "payment"
		} else {
			return s.error(c, flow.ErrInvalidKey, "out of service")
		}

	case hardware.StarKey:
		c.PaymentChannel = nil
		c.ChangeStage <- "summary"

	default:
		log.Debug().Str("stage", "payment_channel").Str("key", key.ToString()).Msg("invalid key")
		return s.error(c, flow.ErrInvalidKey, "invalid key")
	}

	return nil
}
