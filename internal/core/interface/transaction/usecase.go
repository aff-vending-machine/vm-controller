package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/response"
)

type Usecase interface {
	Paid(context.Context, *request.Paid) error
	Cancel(context.Context, *request.Cancel) error
	Get(context.Context) ([]response.Transaction, error)
	Clear(context.Context, *request.Clear) error
}
