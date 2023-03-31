package ui_ws

import (
	"context"
	"time"
)

type OTPRequestedData struct {
	Request   string    `json:"request"`
	OrderID   string    `json:"order_id"`
	Mail      string    `json:"mail"`
	Reference string    `json:"reference"`
	Timestamp time.Time `json:"timestamp"`
}

func (w *wsImpl) SendOTPRequest(ctx context.Context, orderID string, mail string, reference string, timestamp time.Time) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := OTPRequestedData{
		Request:   "otp",
		Reference: reference,
		OrderID:   orderID,
		Mail:      mail,
		Timestamp: timestamp,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "identifiaction",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
