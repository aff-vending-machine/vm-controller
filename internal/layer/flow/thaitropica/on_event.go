package thaitropica

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/rs/zerolog/log"
)

func (uc *Flow) OnEvent(ctx context.Context, event *hardware.Event) error {
	uc.context.UserCtx = ctx

	log.Debug().Str("stage", uc.context.Stage).Interface("event", event).Msg("event occured")

	return uc.stages[uc.context.Stage].OnEvent(uc.context, event)
}
