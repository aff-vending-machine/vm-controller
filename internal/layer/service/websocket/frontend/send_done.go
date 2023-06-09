package frontend

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
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
		Stage: flow.DONE_STAGE,
		Data:  data,
	}

	log.Info().Msg("sending done")
	return w.client.WriteJSON(payload)
}
