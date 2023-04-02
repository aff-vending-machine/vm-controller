package emergency

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc usecase.Display
	reset     int
}

func New(du usecase.Display) *stageImpl {
	return &stageImpl{du, 0}
}
