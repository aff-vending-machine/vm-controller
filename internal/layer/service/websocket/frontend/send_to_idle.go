package frontend

import (
	"context"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (w *wsImpl) SendToIdle(ctx context.Context) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	payload := PayloadModel{
		Code:  200,
		Stage: flow.IDLE_STAGE,
	}

	log.Info().Interface("payload", payload).Msg("sending to idle")
	return w.client.WriteJSON(payload)
}
