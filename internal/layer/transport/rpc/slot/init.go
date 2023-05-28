package slot

import "github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"

type rpcImpl struct {
	usecase slot.Usecase
}

func New(uc slot.Usecase) *rpcImpl {
	return &rpcImpl{uc}
}
