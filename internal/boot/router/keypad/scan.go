package keypad

import "github.com/aff-vending-machine/vm-controller/internal/boot/registry"

func (i *input) Scan(trans registry.KeypadTransport) {
	i.App.SetOnPressed(trans.Keypad.OnPressed)

	go i.App.Listen()
}
