package machine

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
)

type usecaseImpl struct {
	topic       api.Topic
	machineRepo machine.Repository
	queueHw     hardware.Queue
}

func New(tp api.Topic, mr machine.Repository, qh hardware.Queue) *usecaseImpl {
	return &usecaseImpl{tp, mr, qh}
}
