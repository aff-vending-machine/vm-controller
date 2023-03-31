package ui_ws

import (
	"context"
)

type MailRequestData struct {
	Request string `json:"request"`
	OrderID string `json:"order_id"`
}

func (w *wsImpl) SendMailRequest(ctx context.Context, orderID string) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := MailRequestData{
		Request: "mail-request",
		OrderID: orderID,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "identifiaction",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
