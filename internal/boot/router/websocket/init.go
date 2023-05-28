package websocket

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/websocket"
)

type serverImpl struct {
	*websocket.Wrapper
}

func New(client *websocket.Wrapper) *serverImpl {
	return &serverImpl{
		client,
	}
}
