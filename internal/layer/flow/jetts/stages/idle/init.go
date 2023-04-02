package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc   usecase.Display
	machineRepo repository.Machine
}

func New(du usecase.Display, mr repository.Machine) *stageImpl {
	return &stageImpl{du, mr}
}
