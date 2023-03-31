package transaction_usecase

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Paid(ctx context.Context, req *request.Paid) error {
	return nil
}
