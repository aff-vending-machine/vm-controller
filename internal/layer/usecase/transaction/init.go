package transaction

import "github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"

type usecaseImpl struct {
	transactionRepo transaction.Repository
}

func New(r transaction.Repository) *usecaseImpl {
	return &usecaseImpl{r}
}
