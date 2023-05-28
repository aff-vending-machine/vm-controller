package request

import (
	"vm-controller/internal/layer/usecase/machine/model"
)

type Set struct {
	Data model.Machine `json:"data"`
}
