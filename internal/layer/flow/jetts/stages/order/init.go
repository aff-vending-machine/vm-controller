package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	queue      hardware.Queue
	slotRepo   repository.Slot
	frontendWs websocket.Frontend
	slots      []entity.Slot
}

func New(qh hardware.Queue, sr repository.Slot, fw websocket.Frontend) *stageImpl {
	return &stageImpl{qh, sr, fw, make([]entity.Slot, 0)}
}
