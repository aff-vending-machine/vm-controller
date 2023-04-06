package keypad

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/keypad"
)

type input struct {
	*keypad.Wrapper
}

func New(client *keypad.Wrapper) *input {
	return &input{
		client,
	}
}
