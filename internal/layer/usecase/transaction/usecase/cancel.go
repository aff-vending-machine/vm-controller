package transaction_usecase

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Cancel(ctx context.Context, req *request.Cancel) error {
	return nil
}
