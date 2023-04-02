package request

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
)

type Set struct {
	Code     string          `json:"code" query:"code" validate:"required"`
	Stock    int             `json:"stock" validate:"stockValidator"`
	Capacity int             `json:"capacity" validate:"int|min:0"`
	Product  *entity.Product `json:"product,omitempty"`
	IsEnable bool            `json:"is_enable"`
}

func (s *Set) ToFilter() []string {
	return []string{
		fmt.Sprintf("code:=:%s", s.Code),
	}
}

func (s *Set) ToEntity() entity.Slot {
	return entity.Slot{
		Code:     s.Code,
		Stock:    s.Stock,
		Capacity: s.Capacity,
		Product:  s.Product,
		IsEnable: s.IsEnable,
	}
}

func (s *Set) ToJson() map[string]interface{} {
	data := map[string]interface{}{
		"code":      s.Code,
		"stock":     s.Stock,
		"capacity":  s.Capacity,
		"is_enable": s.IsEnable,
	}

	if s.Product != nil {
		data["product_sku"] = s.Product.SKU
		data["product_name"] = s.Product.Name
		data["product_image_url"] = s.Product.ImageURL
		data["product_price"] = s.Product.Price
	}

	return data
}

func (s Set) StockValidator(val int) bool {
	return val >= 0 && val <= s.Capacity
}
