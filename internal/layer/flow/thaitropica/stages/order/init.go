package order

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
)

type Stage struct {
	queue    hardware.Queue
	slotRepo repository.Slot
	ui       ws.UI
	slots    []entity.Slot
}

func New(qh hardware.Queue, s repository.Slot, u ws.UI) *Stage {
	return &Stage{qh, s, u, make([]entity.Slot, 0)}
}
