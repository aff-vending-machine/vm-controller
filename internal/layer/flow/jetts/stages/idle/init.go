package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	machineRepo repository.Machine
	frontendWs  websocket.Frontend
}

func New(mr repository.Machine, fw websocket.Frontend) *stageImpl {
	return &stageImpl{mr, fw}
}
