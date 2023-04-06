package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel"
	payment_channel_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/usecase"
	payment_channel_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/wrapper"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction"
	transaction_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/usecase"
	transaction_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/wrapper"
)

// Usecase layers
type Usecase struct {
	Machine        usecase.Machine
	PaymentChannel interface{ payment_channel.Usecase }
	Slot           usecase.Slot
	Transaction    interface{ transaction.Usecase }
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		Machine: machine.New(
			adapter.Repository.Machine,
			adapter.Hardware.Queue,
		),
		PaymentChannel: payment_channel_wrapper.New(
			payment_channel_usecase.New(
				adapter.Repository.PaymentChannel,
			),
		),
		Slot: slot.New(
			adapter.Repository.Slot,
		),
		Transaction: transaction_wrapper.New(
			transaction_usecase.New(
				adapter.Repository.Transaction,
			),
		),
	}
}
