package identification

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type stageImpl struct {
	mailAPI         api.Mail
	queueHw         hardware.Queue
	customerRepo    repository.Customer
	transactionRepo repository.Transaction
	ui              ws.UI
	stacks          map[string]*OTPStack
}

type OTPStack struct {
	mail      string
	otp       string
	reference string
	timestamp time.Time
}

func New(a api.Mail, q hardware.Queue, c repository.Customer, t repository.Transaction, u ws.UI) *stageImpl {
	return &stageImpl{
		a, q, c, t, u,
		make(map[string]*OTPStack),
	}
}
