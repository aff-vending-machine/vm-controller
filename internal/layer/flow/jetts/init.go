package jetts

import (
	"time"

	"vm-controller/internal/core/flow"
	"vm-controller/internal/core/interface/machine"
	"vm-controller/internal/core/interface/payment_channel"
	"vm-controller/internal/core/interface/slot"
	"vm-controller/internal/core/interface/transaction"
	"vm-controller/internal/layer/flow/jetts/stages"
	"vm-controller/internal/layer/flow/jetts/stages/channel"
	"vm-controller/internal/layer/flow/jetts/stages/emergency"
	"vm-controller/internal/layer/flow/jetts/stages/idle"
	"vm-controller/internal/layer/flow/jetts/stages/order"
	"vm-controller/internal/layer/flow/jetts/stages/payment"
	"vm-controller/internal/layer/flow/jetts/stages/receive"
	"vm-controller/internal/layer/service/api"
	"vm-controller/internal/layer/service/hardware"
	"vm-controller/internal/layer/service/websocket"
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
