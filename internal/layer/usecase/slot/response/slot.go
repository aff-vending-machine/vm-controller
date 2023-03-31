package response

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"

type Slot struct {
	Code     string   `json:"code"`
	Stock    int      `json:"stock"`
	Capacity int      `json:"capacity"`
	Product  *Product `json:"product,omitempty"`
	IsEnable bool     `json:"is_enable"`
}

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func ToModel(e *entity.Slot) *Slot {
	return &Slot{
		Code:     e.Code,
		Product:  SafeToModel(e.Product),
		Stock:    e.Stock,
		Capacity: e.Capacity,
		IsEnable: e.IsEnable,
	}
}

func SafeToModel(e *entity.Product) *Product {
	if e == nil {
		return nil
	}

	return &Product{
		SKU:      e.SKU,
		Name:     e.Name,
		ImageURL: e.ImageURL,
		Price:    e.Price,
	}
}
