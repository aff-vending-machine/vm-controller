package transaction

import "github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"

type rpcImpl struct {
	usecase transaction.Usecase
}

func New(uc transaction.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
