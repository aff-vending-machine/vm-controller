package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/jetts"
)

// Usecase layers
type AppFlow struct {
	Jetts interface{ flow.Jetts }
}

func NewAppFlow(adapter AppDriven) AppFlow {
	return AppFlow{
		jetts.New(
			adapter.API.Ksher,
			adapter.API.Link2500,
			adapter.Asset.Images,
			adapter.Asset.Fonts,
			adapter.Display.LCD,
			adapter.Hardware.Queue,
			adapter.Repository.Machine,
			adapter.Repository.PaymentChannel,
			adapter.Repository.Slot,
			adapter.Repository.Transaction,
		),
	}
}
