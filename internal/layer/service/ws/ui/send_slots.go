package ui_ws

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
)

type OrderInitData struct {
	Slots     []SlotData `json:"slots"`
	FreeSlots []SlotData `json:"free_slots"`
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
	freeSlotData := make([]SlotData, 0)
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

		if s.Price == 0 {
			freeSlotData = append(freeSlotData, s)
		} else {
			slotData = append(slotData, s)
		}
	}

	data := OrderInitData{
		Slots:     slotData,
		FreeSlots: freeSlotData,
	}

	payload := PayloadModel{
		Code:  200,
		Stage: "order",
		Data:  data,
	}

	return w.client.WriteJSON(payload)
}
