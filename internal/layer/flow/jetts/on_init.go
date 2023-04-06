package jetts

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) OnInit(ctx context.Context) {
	uc.context.UserCtx = ctx

	log.Debug().Str("stage", uc.context.Stage).Msg("initial stage")

	if uc.stages[uc.context.Stage] == nil {
		log.Debug().Str("stage", uc.context.Stage).Msg("stage is nil")
		uc.context.ChangeStage <- "order"
	}

	uc.stages[uc.context.Stage].OnInit(uc.context)
}
