package model

import (
	"vm-controller/internal/core/domain/entity"
)

type Slot struct {
	Code     string   `json:"code"`
	Stock    int      `json:"stock"`
	Capacity int      `json:"capacity"`
	Product  *Product `json:"product,omitempty"`
	IsEnable bool     `json:"is_enable"`
}

func (m *Slot) ToEntity() *entity.Slot {
	return &entity.Slot{
		Code:     m.Code,
		Stock:    m.Stock,
		Capacity: m.Capacity,
		Product:  m.Product.ToEntity(),
		IsEnable: m.IsEnable,
	}
}

func (m *Slot) ToJson() map[string]interface{} {
	json := map[string]interface{}{
		"code":      m.Code,
		"stock":     m.Stock,
		"capacity":  m.Capacity,
		"is_enable": m.IsEnable,
	}

	if m.Product != nil {
		json["product_sku"] = m.Product.SKU
		json["product_name"] = m.Product.Name
		json["product_type"] = m.Product.Type
		json["product_image_url"] = m.Product.ImageURL
		json["product_price"] = m.Product.Price
	}

	return json
}

func ToSlot(e *entity.Slot) *Slot {
	return &Slot{
		Code:     e.Code,
		Stock:    e.Stock,
		Capacity: e.Capacity,
		Product:  ToProduct(e.Product),
		IsEnable: e.IsEnable,
	}
}
