package frontend

import (
	"context"

	"vm-controller/internal/core/domain/entity"
)

type PaymentChannelData struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
	Vendor  string `json:"vendor"`
}

func (w *wsImpl) SendPaymentChannel(ctx context.Context, channels []entity.PaymentChannel) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	data := make([]PaymentChannelData, 0)
	for _, channel := range channels {
		data = append(data, PaymentChannelData{
			Name:    channel.Name,
			Channel: channel.Channel,
			Vendor:  channel.Vendor,
		})
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "payment_channel", // old stage for support frontend
		// Stage: flow.CHANNEL_STAGE,
		Data: data,
	}

	return w.client.WriteJSON(payload)
}
