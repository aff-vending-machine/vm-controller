package jetts

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) OnKeyPressed(ctx context.Context, k string) error {
	uc.context.UserCtx = ctx
	uc.context.ClearWatchdog <- true

	log.Debug().Str("stage", uc.context.Stage).Str("k", k).Msg("key pressed")

	if uc.stages[uc.context.Stage] == nil {
		log.Debug().Str("stage", uc.context.Stage).Msg("stage is nil")
		uc.context.ChangeStage <- "order"
		return nil
	}

	return uc.stages[uc.context.Stage].OnKeyPressed(uc.context, k)
}
