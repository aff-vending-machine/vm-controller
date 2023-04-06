package board

import "github.com/aff-vending-machine/vm-controller/internal/layer/flow"

type raspiImpl struct {
	usecase flow.Jetts
}

func New(s flow.Jetts) *raspiImpl {
	return &raspiImpl{s}
}
