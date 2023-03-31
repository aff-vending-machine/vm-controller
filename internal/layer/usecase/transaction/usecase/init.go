package transaction_usecase

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
)

type usecaseImpl struct {
	transactionRepo repository.Transaction
}

func New(r repository.Transaction) *usecaseImpl {
	return &usecaseImpl{r}
}
