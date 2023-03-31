package payment

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type stageImpl struct {
	lugentpay       api.LugentPay
	queueHw         hardware.Queue
	transactionRepo repository.Transaction
	smartedc        serial.SmartEDC
	ui              ws.UI
	delay           time.Duration
	ticker          *time.Ticker
	qrcode          *string
	CancelFn        context.CancelFunc
}

func New(la api.LugentPay, qh hardware.Queue, tr repository.Transaction, se serial.SmartEDC, ui ws.UI) *stageImpl {
	return &stageImpl{
		la, qh, tr, se, ui,
		10 * time.Second,
		nil,
		nil,
		nil,
	}
}
