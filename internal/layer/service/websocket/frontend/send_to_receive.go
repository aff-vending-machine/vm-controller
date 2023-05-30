package frontend

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
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
		Stage: flow.RECEIVE_STAGE,
		Data:  data,
	}

	log.Info().Msg("sending to receive")
	return w.client.WriteJSON(payload)
}
