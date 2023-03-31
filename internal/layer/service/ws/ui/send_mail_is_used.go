package ui_ws

import (
	"context"
)

type MailIsUsedData struct {
	Request string `json:"request"`
	OrderID string `json:"order_id"`
	Mail    string `json:"mail"`
}

func (w *wsImpl) SendMailIsUsed(ctx context.Context, orderID string, mail string) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := MailIsUsedData{
		Request: "mail-is-used",
		OrderID: orderID,
		Mail:    mail,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "identifiaction",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
