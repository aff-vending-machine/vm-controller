package payment_channel

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "payment_channel")
}

func (s *stageImpl) show(c *flow.Ctx, channels []entity.PaymentChannel) {
	log.Info().Str("stage", "payment_channel").Interface("channels", channels).Msg("SLOG: payment channel action")

	s.displayUc.Clear(c.UserCtx)
	s.displayUc.StagePaymentChannel(c.UserCtx, channels)
}

func (s *stageImpl) error(c *flow.Ctx, err error, msg string) error {
	log.Info().Str("stage", "payment_channel").Err(err).Msg("SLOG: payment channel error")

	s.displayUc.Error(c.UserCtx, fmt.Errorf(msg))

	go func() {
		time.Sleep(5 * time.Second)
		if c.Stage == "payment_channel" {
			s.show(c, s.channels)
		}
	}()

	return err
}
