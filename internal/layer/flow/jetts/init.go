package jetts

import (
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/machine"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/emergency"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/idle"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/order"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/payment"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts/stages/receive"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/api"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type Flow struct {
	stages   map[string]stages.Stage
	queueHw  hardware.Queue
	context  *flow.Ctx
	watchdog *time.Ticker
}

func New(
	ka api.Ksher,
	la api.Link2500,
	qh hardware.Queue,
	mr machine.Repository,
	pr payment_channel.Repository,
	sr slot.Repository,
	tr transaction.Repository,
	fw websocket.Frontend,
) *Flow {

	stages := map[string]stages.Stage{
		"idle":      idle.New(mr, fw),
		"order":     order.New(qh, sr, fw),
		"channel":   channel.New(pr, tr, fw),
		"payment":   payment.New(ka, la, qh, tr, fw),
		"receive":   receive.New(ka, la, qh, sr, tr, fw),
		"emergency": emergency.New(fw),
	}

	return &Flow{
		stages:  stages,
		queueHw: qh,
		context: flow.NewContext(),
	}
}
