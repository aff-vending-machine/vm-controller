package jetts

import (
	"context"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) OnInit(ctx context.Context) {
	uc.context.UserCtx = ctx

	log.Debug().Str("stage", string(uc.context.Stage)).Int("quantity", uc.context.Data.TotalQuantity()).Str("payment", uc.context.PaymentChannel.Channel).Msg("initial stage")

	if uc.stages[uc.context.Stage] == nil {
		log.Debug().Str("stage", string(uc.context.Stage)).Msg("stage is nil")
		uc.context.ChangeStage <- flow.ORDER_STAGE
		return
	}

	uc.stages[uc.context.Stage].OnInit(uc.context)
}
