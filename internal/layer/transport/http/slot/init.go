package slot_http

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot"

type restImpl struct {
	usecase slot.Usecase
}

func New(uc slot.Usecase) *restImpl {
	return &restImpl{uc}
}
