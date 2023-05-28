package frontend

import "github.com/aff-vending-machine/vm-controller/internal/layer/flow"

type wsImpl struct {
	usecase flow.Jetts
}

func New(uc flow.Jetts) *wsImpl {
	return &wsImpl{uc}
}
