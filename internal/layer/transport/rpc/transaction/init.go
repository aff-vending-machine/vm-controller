package transaction

import "vm-controller/internal/core/interface/transaction"

type rpcImpl struct {
	usecase transaction.Usecase
}

func New(uc transaction.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
