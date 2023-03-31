package transaction_wrapper

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction"
)

type wrapperImpl struct {
	usecase transaction.Usecase
}

func New(uc transaction.Usecase) *wrapperImpl {
	return &wrapperImpl{uc}
}
