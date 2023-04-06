package keypad

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/stianeikeland/go-rpio/v4"
)

type Wrapper struct {
	*App
}

func New(cfg config.BoardConfig) *Wrapper {
	err := rpio.Open()
	boot.TerminateWhenError(err)
	boot.AddCloseFn(rpio.Close)

	vline := make([]rpio.Pin, 0)
	for _, v := range cfg.KeypadVerticalLine {
		vline = append(vline, rpio.Pin(v))
	}
	hline := make([]rpio.Pin, 0)
	for _, h := range cfg.KeypadHorizontalLine {
		hline = append(hline, rpio.Pin(h))
	}

	app := &App{
		keys: [][]string{
			{"1", "2", "3", "A"},
			{"4", "5", "6", "B"},
			{"7", "8", "9", "C"},
			{"*", "0", "#", "D"},
		},
		vertical:   vline,
		horizontal: hline,
	}

	return &Wrapper{
		app,
	}
}
