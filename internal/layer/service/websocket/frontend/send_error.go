package frontend

import (
	"context"
	"vm-controller/internal/core/flow"
)

type ErrorData struct {
	Message string `json:"message"`
}

func (w *wsImpl) SendError(ctx context.Context, stage flow.Stage, message string) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := ErrorData{
		Message: message,
	}

	payload := PayloadModel{
		Code:  500,
		Stage: string(stage),
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
