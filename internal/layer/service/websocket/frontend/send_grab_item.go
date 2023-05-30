package frontend

import (
	"context"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

type GrabItemData struct {
	Message string `json:"message"`
}

func (w *wsImpl) SendGrabItem(ctx context.Context, stage flow.Stage, message string) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := GrabItemData{
		Message: message,
	}
	s := string(stage)
	if stage == flow.CHANNEL_STAGE {
		s = "payment_channel"
	}

	payload := PayloadModel{
		Code:  501,
		Stage: s,
		Data:  data,
	}

	log.Info().Interface("payload", payload).Msg("sending grab item")
	return w.client.WriteJSON(payload)
}
