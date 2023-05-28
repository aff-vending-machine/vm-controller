package request

import (
	"vm-controller/internal/layer/usecase/payment_channel/model"
)

type Set struct {
	Data []model.PaymentChannel `json:"data"`
}
