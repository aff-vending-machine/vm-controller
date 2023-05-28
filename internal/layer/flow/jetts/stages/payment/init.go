package payment

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	ksher           api.Ksher
	link2500        api.Link2500
	queue           hardware.Queue
	transactionRepo transaction.Repository
	frontendWs      websocket.Frontend
	delay           time.Duration
	ticker          *time.Ticker
	qrcode          *string
	CancelFn        context.CancelFunc
}

func New(ka api.Ksher, la api.Link2500, qh hardware.Queue, tr transaction.Repository, fw websocket.Frontend) *stageImpl {
	return &stageImpl{
		ka, la, qh, tr, fw,
		10 * time.Second,
		nil,
		nil,
		nil,
	}
}
