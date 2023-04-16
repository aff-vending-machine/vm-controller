package emergency

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/websocket"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase"
)

type stageImpl struct {
	displayUc  usecase.Screen
	frontendWs websocket.Frontend
	reset      int
}

func New(du usecase.Screen, fw websocket.Frontend) *stageImpl {
	return &stageImpl{du, fw, 0}
}
