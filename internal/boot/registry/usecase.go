package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine"
	machine_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/usecase"
	machine_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/wrapper"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel"
	payment_channel_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/usecase"
	payment_channel_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/wrapper"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"
	slot_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/usecase"
	slot_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/wrapper"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction"
	transaction_usecase "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/usecase"
	transaction_wrapper "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/wrapper"
)

// Usecase layers
type Usecase struct {
	Machine        interface{ machine.Usecase }
	PaymentChannel interface{ payment_channel.Usecase }
	Slot           interface{ slot.Usecase }
	Transaction    interface{ transaction.Usecase }
}

func NewUsecase(adapter Service) Usecase {
	return Usecase{
		Machine: machine_wrapper.New(
			machine_usecase.New(
				adapter.Repository.Machine,
				adapter.Hardware.Queue,
			),
		),
		PaymentChannel: payment_channel_wrapper.New(
			payment_channel_usecase.New(
				adapter.Repository.PaymentChannel,
			),
		),
		Slot: slot_wrapper.New(
			slot_usecase.New(
				adapter.Repository.Slot,
			),
		),
		Transaction: transaction_wrapper.New(
			transaction_usecase.New(
				adapter.Repository.Transaction,
			),
		),
	}
}
