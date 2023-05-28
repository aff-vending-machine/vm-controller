package request

import (
	"vm-controller/internal/layer/usecase/product/model"
)

type Set struct {
	Data model.Product `json:"data"`
}
