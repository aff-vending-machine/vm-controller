package payment

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc       usecase.Screen
	ksher           api.Ksher
	link2500        api.Link2500
	queue           hardware.Queue
	transactionRepo repository.Transaction
	delay           time.Duration
	ticker          *time.Ticker
	qrcode          *string
	CancelFn        context.CancelFunc
}

func New(du usecase.Screen, ka api.Ksher, la api.Link2500, qh hardware.Queue, tr repository.Transaction) *stageImpl {
	return &stageImpl{
		du, ka, la, qh, tr,
		10 * time.Second,
		nil,
		nil,
		nil,
	}
}
