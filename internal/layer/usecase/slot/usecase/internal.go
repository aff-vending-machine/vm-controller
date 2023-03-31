package slot_usecase

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/response"
)

func toResponseList(slots []entity.Slot) []response.Slot {
	res := make([]response.Slot, 0)
	for _, slot := range slots {
		var product response.Product
		if slot.Product != nil {
			product = response.Product{
				SKU:      slot.Product.SKU,
				Name:     slot.Product.Name,
				ImageURL: slot.Product.ImageURL,
				Price:    slot.Product.Price,
			}
		}

		res = append(res, response.Slot{
			Code:     slot.Code,
			Product:  &product,
			Stock:    slot.Stock,
			Capacity: slot.Capacity,
			IsEnable: slot.IsEnable,
		})
	}

	return res
}
