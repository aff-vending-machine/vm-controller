package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/request"
)

func (uc *usecaseImpl) Clear(ctx context.Context, req *request.Clear) error {
	_, err := uc.transactionRepo.DeleteMany(ctx, req.ToFilter())
	if err != nil {
		return err
	}

	return nil
}
