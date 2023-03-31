package ui_ws

import (
	"context"
)

type IdentifiedData struct {
	OrderID  string  `json:"order_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (w *wsImpl) SendIdentified(ctx context.Context, orderID string, qty int, price float64) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := PaidData{
		OrderID:  orderID,
		Quantity: qty,
		Price:    price,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "identifiaction",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
