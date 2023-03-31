package machine_wrapper

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine"
)

type wrapperImpl struct {
	usecase machine.Usecase
}

func New(uc machine.Usecase) *wrapperImpl {
	return &wrapperImpl{uc}
}
