package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	ksher           api.Ksher
	link2500        api.Link2500
	queue           hardware.Queue
	slotRepo        slot.Repository
	transactionRepo transaction.Repository
	frontendWs      websocket.Frontend
	status          int
	polling         bool
}

const (
	WAIT   = 0
	CANCEL = 1
	DONE   = 2
	E0     = 0xE0
	E1     = 0xE1
	E2     = 0xE2
)

func New(ka api.Ksher, la api.Link2500, qh hardware.Queue, sr slot.Repository, tr transaction.Repository, fw websocket.Frontend) *stageImpl {
	return &stageImpl{ka, la, qh, sr, tr, fw, 0, false}
}
