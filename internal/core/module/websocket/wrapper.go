package websocket

import (
	"fmt"
	"net/http"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/gorilla/websocket"
)

type Wrapper struct {
	App  websocket.Upgrader
	Port string
}

func New(conf config.WebSocketConfig) *Wrapper {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Wrapper{
		App:  upgrader,
		Port: fmt.Sprintf(":%s", conf.Port),
	}
}
