package machine

import "vm-controller/internal/core/interface/machine"

type rpcImpl struct {
	usecase machine.Usecase
}

func New(uc machine.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
