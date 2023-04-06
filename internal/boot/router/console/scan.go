package console

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/rs/zerolog/log"
)

func (i *input) Scan(svr registry.KeypadTransport) {
	i.App.SetOnPressed(svr.Keypad.OnPressed)

	go i.App.Listen()

	log.Info().Msg("scan keyboard")
}
