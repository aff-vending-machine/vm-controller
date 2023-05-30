package frontend

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
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
		Stage: flow.RECEIVE_STAGE,
		Data:  data,
	}

	log.Info().Interface("payload", payload).Msg("sending received item")
	return w.client.WriteJSON(payload)
}
