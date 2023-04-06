package machine

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
)

type usecaseImpl struct {
	machineRepo repository.Machine
	queueHw     hardware.Queue
}

func New(mr repository.Machine, qh hardware.Queue) *usecaseImpl {
	return &usecaseImpl{mr, qh}
}
