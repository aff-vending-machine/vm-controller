package frontend

import (
	"context"
	"vm-controller/internal/core/flow"
)

type QRCodeData struct {
	OrderID  string  `json:"order_id"`
	QRCode   string  `json:"qrcode"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (w *wsImpl) SendQRCode(ctx context.Context, orderID string, qrcode string, qty int, price float64) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := QRCodeData{
		OrderID:  orderID,
		QRCode:   qrcode,
		Quantity: qty,
		Price:    price,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: flow.PAYMENT_STAGE,
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
