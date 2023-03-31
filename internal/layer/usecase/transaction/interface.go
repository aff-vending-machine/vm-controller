package transaction

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/request"
)

type Usecase interface {
	Paid(context.Context, *request.Paid) error
	Cancel(context.Context, *request.Cancel) error
}
