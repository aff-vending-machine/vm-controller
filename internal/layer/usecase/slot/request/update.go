package request

import (
	"encoding/json"
	"fmt"
)

type Update struct {
	ID              uint     `json:"-" query:"id" validate:"required"`
	Code            *string  `json:"code,omitempty"`
	Stock           *int     `json:"stock,omitempty"`
	Capacity        *int     `json:"capacity,omitempty"`
	IsEnable        *bool    `json:"is_enable,omitempty"`
	ProductSKU      *string  `json:"product_sku,omitempty"`
	ProductName     *string  `json:"product_name,omitempty"`
	ProductType     *string  `json:"product_type,omitempty"`
	ProductImageURL *string  `json:"product_image_url,omitempty"`
	ProductPrice    *float64 `json:"product_price,omitempty"`
}

func (s *Update) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", s.ID),
	}
}

func (s *Update) ToJson() map[string]interface{} {
	var res map[string]interface{}
	b, _ := json.Marshal(s)
	json.Unmarshal(b, &res)
	return res
}
