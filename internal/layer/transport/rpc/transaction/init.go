package transaction

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type rpcImpl struct {
	usecase usecase.Transaction
}

func New(uc usecase.Transaction) *rpcImpl {
	return &rpcImpl{uc}
}
