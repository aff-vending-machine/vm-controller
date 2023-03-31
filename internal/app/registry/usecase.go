package registry

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine"
	machine_usecase "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/usecase"
	machine_wrapper "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/wrapper"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel"
	payment_channel_usecase "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel/usecase"
	payment_channel_wrapper "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel/wrapper"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot"
	slot_usecase "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/usecase"
	slot_wrapper "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/wrapper"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction"
	transaction_usecase "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/usecase"
	transaction_wrapper "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/wrapper"
)

// Usecase layers
type AppUsecase struct {
	Machine        interface{ machine.Usecase }
	PaymentChannel interface{ payment_channel.Usecase }
	Slot           interface{ slot.Usecase }
	Transaction    interface{ transaction.Usecase }
}

func NewAppUsecase(adapter AppDriven) AppUsecase {
	return AppUsecase{
		machine_wrapper.New(
			machine_usecase.New(
				adapter.Repository.Machine,
				adapter.Hardware.Queue,
			),
		),
		payment_channel_wrapper.New(
			payment_channel_usecase.New(
				adapter.Repository.PaymentChannel,
			),
		),
		slot_wrapper.New(
			slot_usecase.New(
				adapter.Repository.Slot,
			),
		),
		transaction_wrapper.New(
			transaction_usecase.New(
				adapter.Repository.Transaction,
			),
		),
	}
}
