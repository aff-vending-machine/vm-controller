package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc       usecase.Screen
	ksher           api.Ksher
	link2500        api.Link2500
	queue           hardware.Queue
	slotRepo        repository.Slot
	transactionRepo repository.Transaction
	frontendWs      websocket.Frontend
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

func New(du usecase.Screen, ka api.Ksher, la api.Link2500, qh hardware.Queue, sr repository.Slot, tr repository.Transaction, fw websocket.Frontend) *stageImpl {
	return &stageImpl{du, ka, la, qh, sr, tr, fw, false, 0, 0, 0}
}
