package frontend

import "vm-controller/internal/layer/flow"

type wsImpl struct {
	usecase flow.Jetts
}

func New(uc flow.Jetts) *wsImpl {
	return &wsImpl{uc}
}
