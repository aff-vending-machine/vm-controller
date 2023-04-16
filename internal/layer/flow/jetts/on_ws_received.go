package jetts

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type WSReceived struct {
	Action string `json:"action"`
}

func (uc *Flow) OnWSReceived(ctx context.Context, data []byte) error {
	uc.context.UserCtx = ctx

	log.Debug().Str("stage", uc.context.Stage).Bytes("data", data).Msg("ws received")

	var req WSReceived
	err := json.Unmarshal(data, &req)
	if err != nil {
		return errors.Wrap(err, "convert to struct failed")
	}

	if uc.stages[uc.context.Stage] == nil {
		log.Debug().Str("stage", uc.context.Stage).Msg("stage is nil")
		uc.context.ChangeStage <- "order"
		return nil
	}

	return uc.stages[uc.context.Stage].OnWSReceived(uc.context, data)
}
