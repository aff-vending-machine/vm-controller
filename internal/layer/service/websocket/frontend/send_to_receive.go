package frontend

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

type ToReceiveData struct {
	OrderID       string          `json:"order_id"`
	Cart          []hardware.Item `json:"cart"`
	TotalPrice    float64         `json:"total_price"`
	TotalQuantity int             `json:"total_quantity"`
}

func (w *wsImpl) SendToReceive(ctx context.Context, orderID string, cart []hardware.Item) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	total := calculateTotal(cart)

	data := ToReceiveData{
		OrderID:       orderID,
		Cart:          cart,
		TotalPrice:    total.Price,
		TotalQuantity: total.Quantity,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "receive",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
