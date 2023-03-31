package machine_usecase

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
)

type usecaseImpl struct {
	machineRepo repository.Machine
	queueHw     hardware.Queue
}

func New(mr repository.Machine, qh hardware.Queue) *usecaseImpl {
	return &usecaseImpl{mr, qh}
}
