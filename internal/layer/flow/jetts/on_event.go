package jetts

import (
	"context"

	"vm-controller/internal/core/domain/hardware"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) OnEvent(ctx context.Context, event *hardware.Event) error {
	uc.context.UserCtx = ctx

	log.Debug().Str("stage", string(uc.context.Stage)).Interface("event", event).Msg("event occured")

	return uc.stages[uc.context.Stage].OnEvent(uc.context, event)
}
