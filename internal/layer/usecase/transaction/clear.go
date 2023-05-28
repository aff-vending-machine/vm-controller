package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
)

func (uc *usecaseImpl) Clear(ctx context.Context, req *request.Clear) error {
	_, err := uc.transactionRepo.Delete(ctx, db.NewQuery().AddWhere("id IN ?", req.Query.IDs))
	if err != nil {
		return err
	}

	return nil
}
