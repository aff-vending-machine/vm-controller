package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/request"
	"vm-controller/internal/layer/usecase/payment_channel/response"
)

type Usecase interface {
	Create(context.Context, *request.Create) error
	Get(context.Context, *request.Get) (*response.PaymentChannel, error)
}
