package frontend

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
)

type OrderData struct {
	Cart          []hardware.Item `json:"cart"`
	TotalPrice    float64         `json:"total_price"`
	TotalQuantity int             `json:"total_quantity"`
}

func (w *wsImpl) SendCart(ctx context.Context, cart []hardware.Item) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	total := calculateTotal(cart)

	data := OrderData{
		Cart:          cart,
		TotalPrice:    total.Price,
		TotalQuantity: total.Quantity,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "order",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
