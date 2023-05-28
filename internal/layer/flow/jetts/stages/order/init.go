package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	queue      hardware.Queue
	slotRepo   slot.Repository
	frontendWs websocket.Frontend
	slots      []entity.Slot
}

func New(qh hardware.Queue, sr slot.Repository, fw websocket.Frontend) *stageImpl {
	return &stageImpl{qh, sr, fw, make([]entity.Slot, 0)}
}
