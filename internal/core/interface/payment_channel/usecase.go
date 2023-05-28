package payment_channel

import (
	"context"

	"vm-controller/internal/layer/usecase/payment_channel/request"
	"vm-controller/internal/layer/usecase/payment_channel/response"
)

type Usecase interface {
	Create(context.Context, *request.Create) error
	GetOne(context.Context, *request.GetOne) (*response.PaymentChannel, error)
	Get(context.Context) ([]response.PaymentChannel, error)
	Set(context.Context, *request.Set) error
}
