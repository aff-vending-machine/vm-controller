package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow"
	"github.com/aff-vending-machine/vm-controller/internal/layer/flow/thaitropica"
)

// Usecase layers
type AppFlow struct {
	ThaiTropica interface{ flow.Jetts }
}

func NewAppFlow(adapter AppDriven) AppFlow {
	return AppFlow{
		thaitropica.New(
			adapter.API.LugentPay,
			adapter.API.Mail,
			adapter.Hardware.Queue,
			adapter.Repository.Customer,
			adapter.Repository.Machine,
			adapter.Repository.PaymentChannel,
			adapter.Repository.Slot,
			adapter.Repository.Transaction,
		),
	}
}
