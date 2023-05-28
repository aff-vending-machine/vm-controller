package machine

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type rpcImpl struct {
	usecase usecase.Machine
}

func New(uc usecase.Machine) *rpcImpl {
	return &rpcImpl{uc}
}
