package request

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/model"
)

type Create struct {
	Code     string         `json:"code" validate:"required"`
	Stock    int            `json:"stock" validate:"required"`
	Capacity int            `json:"capacity" validate:"required"`
	Product  *model.Product `json:"product,omitempty"`
	IsEnable bool           `json:"is_enable" validate:"required"`
}

func (s *Create) ToEntity() *entity.Slot {
	return &entity.Slot{
		Code:     s.Code,
		Stock:    s.Stock,
		Capacity: s.Capacity,
		Product:  s.Product.ToEntity(),
		IsEnable: s.IsEnable,
	}
}
