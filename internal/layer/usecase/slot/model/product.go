package model

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
)

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func (s *Product) ToEntity() *entity.Product {
	return &entity.Product{
		SKU:      s.SKU,
		Name:     s.Name,
		Type:     s.Type,
		ImageURL: s.ImageURL,
		Price:    s.Price,
	}
}

func ToProduct(e *entity.Product) *Product {
	if e == nil {
		return nil
	}

	return &Product{
		SKU:      e.SKU,
		Name:     e.Name,
		Type:     e.Type,
		ImageURL: e.ImageURL,
		Price:    e.Price,
	}
}
