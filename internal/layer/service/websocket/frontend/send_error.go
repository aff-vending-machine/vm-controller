package frontend

import "context"

type ErrorData struct {
	Message string `json:"message"`
}

func (w *wsImpl) SendError(ctx context.Context, stage string, message string) error {
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
		Stage: stage,
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
