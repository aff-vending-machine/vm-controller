package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Paid(ctx context.Context, req *request.Paid) error {
	return nil
}
