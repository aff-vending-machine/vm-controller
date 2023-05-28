package frontend

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

type ReceivedItemData struct {
	OrderID       string          `json:"order_id"`
	Cart          []hardware.Item `json:"cart"`
	ReceivedItem  hardware.Item   `json:"received_item"`
	TotalPrice    float64         `json:"total_price"`
	TotalQuantity int             `json:"total_quantity"`
	TotalReceived int             `json:"total_received"`
}

func (w *wsImpl) SendReceivedItem(ctx context.Context, orderID string, cart []hardware.Item, item hardware.Item) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	total := calculateTotal(cart)

	data := ReceivedItemData{
		OrderID:       orderID,
		ReceivedItem:  item,
		TotalQuantity: total.Quantity,
		TotalReceived: total.Received,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "receive",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
