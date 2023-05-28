package frontend

import (
	"sync"

	"github.com/gorilla/websocket"
)

type wsImpl struct {
	client *websocket.Conn
	mu     *sync.Mutex
}

func New() *wsImpl {
	return &wsImpl{
		mu: &sync.Mutex{},
	}
}
