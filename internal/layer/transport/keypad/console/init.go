package console

import "github.com/aff-vending-machine/vm-controller/internal/layer/flow"

type hostImpl struct {
	usecase flow.Jetts
}

func New(uc flow.Jetts) *hostImpl {
	return &hostImpl{uc}
}
