package request

import (
	"fmt"
)

type SetStock struct {
	ID       uint `json:"id" query:"id" validate:"required"`
	Stock    int  `json:"stock" validate:"required"`
	Capacity int  `json:"capacity,omitempty"`
}

func (s *SetStock) ToFilter() []string {
	return []string{
		fmt.Sprintf("id:=:%d", s.ID),
	}
}

func (s *SetStock) ToJson() map[string]interface{} {
	data := map[string]interface{}{
		"stock": s.Stock,
	}

	if s.Capacity > 0 {
		data["capacity"] = s.Capacity
	}

	return data
}
