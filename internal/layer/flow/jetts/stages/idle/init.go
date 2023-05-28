package idle

import (
	"vm-controller/internal/core/interface/machine"
	"vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	machineRepo machine.Repository
	frontendWs  websocket.Frontend
}

func New(mr machine.Repository, fw websocket.Frontend) *stageImpl {
	return &stageImpl{mr, fw}
}
