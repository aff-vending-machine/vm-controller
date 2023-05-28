package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts"
)

func NewFlow(adapter modules.Service) *jetts.Flow {
	return jetts.New(
		adapter.API.Ksher,
		adapter.API.Link2500,
		adapter.Hardware.Queue,
		adapter.Repository.Machine,
		adapter.Repository.PaymentChannel,
		adapter.Repository.Slot,
		adapter.Repository.Transaction,
		adapter.WebSocket.Frontend,
	)
}
