package ws

import (
	"fmt"
	"net/http"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/gorilla/websocket"
)

type serverImpl struct {
	app  websocket.Upgrader
	port string
}

func New(conf config.WebSocketConfig) *serverImpl {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &serverImpl{
		app:  upgrader,
		port: fmt.Sprintf(":%s", conf.Port),
	}
}
