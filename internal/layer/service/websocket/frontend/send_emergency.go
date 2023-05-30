package frontend

import (
	"context"
)

type EmergencyData struct {
	Message string `json:"message"`
}

func (w *wsImpl) SendEmergency(ctx context.Context, err error) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := ErrorData{
		Message: err.Error(),
	}

	payload := PayloadModel{
		Code: 500,
		// Stage: flow.EMERGENCY_STAGE,
		Stage: "error",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
