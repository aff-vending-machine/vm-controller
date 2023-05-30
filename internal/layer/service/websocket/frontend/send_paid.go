package frontend

import (
	"context"
	"vm-controller/internal/core/flow"
)

type PaidData struct {
	OrderID  string  `json:"order_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Paid     bool    `json:"paid"`
}

func (w *wsImpl) SendPaid(ctx context.Context, orderID string, qty int, price float64) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := PaidData{
		OrderID:  orderID,
		Quantity: qty,
		Price:    price,
		Paid:     true,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: flow.PAYMENT_STAGE,
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
