package frontend

import (
	"context"

	"github.com/gorilla/websocket"
)

func (w *wsImpl) SetConnection(ctx context.Context, client *websocket.Conn) {
	w.client = client
}
