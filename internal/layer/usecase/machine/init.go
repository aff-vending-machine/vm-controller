package machine

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
)

type usecaseImpl struct {
	topic       api.Topic
	machineRepo repository.Machine
	queueHw     hardware.Queue
}

func New(tp api.Topic, mr repository.Machine, qh hardware.Queue) *usecaseImpl {
	return &usecaseImpl{tp, mr, qh}
}
