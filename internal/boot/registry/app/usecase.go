package app

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction"
)

func NewUsecase(adapter registry.Service) registry.Usecase {
	return registry.Usecase{
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
