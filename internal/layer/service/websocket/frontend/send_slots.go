package frontend

import (
	"context"

	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/flow"
)

type OrderInitData struct {
	Slots []SlotData `json:"slots"`
}

type SlotData struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
	Stock    int     `json:"stock"`
	IsEnable bool    `json:"is_enable"`
}

func (w *wsImpl) SendSlots(ctx context.Context, slots []entity.Slot) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	slotData := make([]SlotData, 0)
	for _, slot := range slots {
		s := SlotData{
			Code:     slot.Code,
			Stock:    slot.Stock,
			IsEnable: slot.IsEnable && slot.Product != nil,
		}

		if slot.Product != nil {
			s.Name = slot.Product.Name
			s.ImageURL = slot.Product.ImageURL
			s.Price = slot.Product.Price
		}

		slotData = append(slotData, s)
	}

	data := OrderInitData{
		Slots: slotData,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: flow.ORDER_STAGE,
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
