package registry

import (
	"vm-controller/internal/boot/modules"
	"vm-controller/internal/layer/usecase/machine"
	"vm-controller/internal/layer/usecase/payment_channel"
	"vm-controller/internal/layer/usecase/slot"
	"vm-controller/internal/layer/usecase/transaction"
)

func NewUsecase(adapter modules.Service) modules.Usecase {
	return modules.Usecase{
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
