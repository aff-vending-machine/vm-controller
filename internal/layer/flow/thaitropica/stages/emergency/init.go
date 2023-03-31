package emergency

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"

type Stage struct {
	ui ws.UI
}

func New(w ws.UI) *Stage {
	return &Stage{w}
}
