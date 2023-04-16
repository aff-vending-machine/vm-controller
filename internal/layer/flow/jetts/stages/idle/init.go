package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc   usecase.Screen
	machineRepo repository.Machine
	frontendWs  websocket.Frontend
}

func New(du usecase.Screen, mr repository.Machine, fw websocket.Frontend) *stageImpl {
	return &stageImpl{du, mr, fw}
}
