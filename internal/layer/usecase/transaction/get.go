package transaction

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/response"
)

func (uc *usecaseImpl) Get(ctx context.Context) ([]response.Transaction, error) {
	transactions, err := uc.transactionRepo.FindMany(ctx, []string{})
	if err != nil {
		return nil, err
	}

	return response.ToTransactionList(transactions), nil
}
