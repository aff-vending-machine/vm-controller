package payment_channel

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/response"
)

type Usecase interface {
	Create(context.Context, *request.Create) error
	Get(context.Context, []string) (*response.PaymentChannel, error)
}
