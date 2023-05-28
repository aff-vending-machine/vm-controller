package websocket

import (
	"net/http"

	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/rs/zerolog/log"
)

func (s *serverImpl) Serve(driven modules.WebSocketService, driver modules.WebSocketTransport) {
	http.HandleFunc("/ws", s.listen(driven.Frontend, driver.Frontend))

	log.Debug().Str("port", s.Port).Msg("web socket server listening...")
	go http.ListenAndServe(s.Port, nil)
}
