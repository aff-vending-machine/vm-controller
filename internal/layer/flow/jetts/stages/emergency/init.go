package emergency

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc usecase.Screen
	reset     int
}

func New(du usecase.Screen) *stageImpl {
	return &stageImpl{du, 0}
}
