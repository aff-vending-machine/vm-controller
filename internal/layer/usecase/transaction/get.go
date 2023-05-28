package transaction

import (
	"context"

	"vm-controller/internal/layer/usecase/transaction/response"
	"vm-controller/pkg/helpers/db"
)

func (uc *usecaseImpl) Get(ctx context.Context) ([]response.Transaction, error) {
	transactions, err := uc.transactionRepo.FindMany(ctx, db.NewQuery())
	if err != nil {
		return nil, err
	}

	return response.ToTransactionList(transactions), nil
}
