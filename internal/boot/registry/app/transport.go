package app

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/keypad/board"
)

func NewTransport(uc registry.Usecase, fw registry.Flow) registry.Transport {
	return registry.Transport{
		Keypad: registry.KeypadTransport{
			Keypad: board.New(fw.Jetts),
		},
	}
}
