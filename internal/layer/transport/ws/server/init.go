package server_ws

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow"

type wsImpl struct {
	usecase flow.ThaiTropica
}

func New(uc flow.ThaiTropica) *wsImpl {
	return &wsImpl{uc}
}
