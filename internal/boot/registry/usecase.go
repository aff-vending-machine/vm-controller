package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction"
)

// Usecase layers
type Usecase struct {
	Machine        usecase.Machine
	PaymentChannel usecase.PaymentChannel
	Slot           usecase.Slot
	Transaction    usecase.Transaction
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		Machine: machine.New(
			adapter.API.Topic,
			adapter.Repository.Machine,
			adapter.Hardware.Queue,
		),
		PaymentChannel: payment_channel.New(
			adapter.Repository.PaymentChannel,
		),
		Slot: slot.New(
			adapter.Repository.Slot,
		),
		Transaction: transaction.New(
			adapter.Repository.Transaction,
		),
	}
}
