package model

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
)

type Slot struct {
	Code     string   `json:"code"`
	Stock    int      `json:"stock"`
	Capacity int      `json:"capacity"`
	Product  *Product `json:"product,omitempty"`
	IsEnable bool     `json:"is_enable"`
}

func (s *Slot) ToEntity() *entity.Slot {
	return &entity.Slot{
		Code:     s.Code,
		Stock:    s.Stock,
		Capacity: s.Capacity,
		Product:  s.Product.ToEntity(),
		IsEnable: s.IsEnable,
	}
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
