package emergency

import (
	"vm-controller/internal/layer/service/websocket"
)

type stageImpl struct {
	frontendWs websocket.Frontend
	reset      int
}

func New(fw websocket.Frontend) *stageImpl {
	return &stageImpl{fw, 0}
}
