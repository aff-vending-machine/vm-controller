package ui_ws

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
)

type DoneData struct {
	OrderID       string          `json:"order_id"`
	Cart          []hardware.Item `json:"cart"`
	TotalPrice    float64         `json:"total_price"`
	TotalQuantity int             `json:"total_quantity"`
	TotalReceived int             `json:"total_received"`
}

func (w *wsImpl) SendDone(ctx context.Context, orderID string, cart []hardware.Item) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	total := calculateTotal(cart)

	data := DoneData{
		OrderID:       orderID,
		Cart:          cart,
		TotalPrice:    total.Price,
		TotalQuantity: total.Quantity,
		TotalReceived: total.Received,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "done",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
