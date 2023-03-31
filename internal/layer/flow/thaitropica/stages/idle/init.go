package idle

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type Stage struct {
	machineRepo repository.Machine
	ui          ws.UI
}

func New(m repository.Machine, u ws.UI) *Stage {
	return &Stage{m, u}
}
