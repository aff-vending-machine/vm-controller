package app

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/keypad/board"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/machine"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc/transaction"
)

func NewTransport(uc registry.Usecase, fw registry.Flow) registry.Transport {
	return registry.Transport{
		Keypad: registry.KeypadTransport{
			Keypad: board.New(fw.Jetts),
		},
		RPC: registry.RPCTransport{
			Machine:     machine.New(uc.Machine),
			Slot:        slot.New(uc.Slot),
			Transaction: transaction.New(uc.Transaction),
		},
	}
}
