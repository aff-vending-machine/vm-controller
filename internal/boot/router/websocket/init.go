package websocket

import "vm-controller/internal/core/infra/network/websocket"

type serverImpl struct {
	*websocket.Wrapper
}

func New(client *websocket.Wrapper) *serverImpl {
	return &serverImpl{
		client,
	}
}
