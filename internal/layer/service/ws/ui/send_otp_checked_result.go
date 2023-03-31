package ui_ws

import (
	"context"
)

type ResultData struct {
	Result bool `json:"result"`
}

func (w *wsImpl) SendOTPCheckedResult(ctx context.Context, result bool) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	payload := PayloadModel{
		Code:  200,
		Stage: "identifiaction",
		Data:  ResultData{result},
	}

	return w.client.WriteJSON(payload)
}
