package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	model "github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc   usecase.Screen
	queue       hardware.Queue
	slotRepo    repository.Slot
	frontendWs  websocket.Frontend
	slot        *entity.Slot
	slots       []entity.Slot
	step        int
	pendingItem model.Item
}

func New(du usecase.Screen, qh hardware.Queue, sr repository.Slot, fw websocket.Frontend) *stageImpl {
	return &stageImpl{du, qh, sr, fw, nil, make([]entity.Slot, 0), 0, model.Item{}}
}
