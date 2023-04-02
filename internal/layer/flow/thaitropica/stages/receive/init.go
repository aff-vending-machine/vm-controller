package receive

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type stageImpl struct {
	lugentpay       api.LugentPay
	queue           hardware.Queue
	customerRepo    repository.Customer
	slotRepo        repository.Slot
	transactionRepo repository.Transaction
	smartedc        serial.SmartEDC
	ui              ws.UI
	status          int
	polling         bool
}

const (
	WAIT   = 0
	DONE   = 1
	CANCEL = 2
	E0     = 0xE0
	E1     = 0xE1
	E2     = 0xE2
)

func New(la api.LugentPay, qh hardware.Queue, cr repository.Customer, sr repository.Slot, tr repository.Transaction, ss serial.SmartEDC, uw ws.UI) *stageImpl {
	return &stageImpl{la, qh, cr, sr, tr, ss, uw, 0, false}
}
