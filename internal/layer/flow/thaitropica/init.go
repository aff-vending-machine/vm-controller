package thaitropica

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/emergency"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/identification"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/idle"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/order"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/payment"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/payment_channel"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/flow/thaitropica/stages/receive"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/api"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/serial"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

type Flow struct {
	stages  map[string]stages.Stage
	queueHw hardware.Queue
	context *flow.Ctx
}

func New(
	la api.LugentPay,
	sa api.Mail,
	qh hardware.Queue,
	cr repository.Customer,
	mr repository.Machine,
	pr repository.PaymentChannel,
	sr repository.Slot,
	tr repository.Transaction,
	ss serial.SmartEDC,
	uw ws.UI,
) *Flow {
	stages := map[string]stages.Stage{
		"idle":            idle.New(mr, uw),
		"order":           order.New(qh, sr, uw),
		"identification":  identification.New(sa, qh, cr, tr, uw),
		"payment_channel": payment_channel.New(pr, uw),
		"payment":         payment.New(la, qh, tr, ss, uw),
		"receive":         receive.New(la, qh, cr, sr, tr, ss, uw),
		"emergency":       emergency.New(uw),
	}

	return &Flow{
		stages:  stages,
		queueHw: qh,
		context: flow.NewContext(),
	}
}
