package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	ksher           api.Ksher
	link2500        api.Link2500
	displayUc       usecase.Display
	queue           hardware.Queue
	slotRepo        repository.Slot
	transactionRepo repository.Transaction
	polling         bool
	status          int
	starcode        int
	sharpcode       int
}

const (
	WAIT   = 0
	CANCEL = 1
	DONE   = 2
	E0     = 0xE0
	E1     = 0xE1
	E2     = 0xE2
)

func New(ka api.Ksher, la api.Link2500, du usecase.Display, qh hardware.Queue, sr repository.Slot, tr repository.Transaction) *stageImpl {
	return &stageImpl{ka, la, du, qh, sr, tr, false, 0, 0, 0}
}
