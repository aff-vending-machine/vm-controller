package console

import "github.com/aff-vending-machine/vm-controller/internal/core/module/console"

type input struct {
	*console.Wrapper
}

func New(client *console.Wrapper) *input {
	return &input{
		client,
	}
}
