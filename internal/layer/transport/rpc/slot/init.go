package slot

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type rpcImpl struct {
	usecase usecase.Slot
}

func New(uc usecase.Slot) *rpcImpl {
	return &rpcImpl{uc}
}
