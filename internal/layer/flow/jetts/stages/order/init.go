package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	model "github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc   usecase.Screen
	queue       hardware.Queue
	slotRepo    repository.Slot
	slot        *entity.Slot
	step        int
	pendingItem model.Item
}

func New(du usecase.Screen, qh hardware.Queue, sr repository.Slot) *stageImpl {
	return &stageImpl{du, qh, sr, nil, 0, model.Item{}}
}
