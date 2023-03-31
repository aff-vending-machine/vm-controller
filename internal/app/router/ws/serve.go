package ws

import (
	"net/http"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/app/registry"
	"github.com/rs/zerolog/log"
)

func (s *serverImpl) Serve(driven registry.WebSocketDriven, driver registry.WebSocketDriver) {
	http.HandleFunc("/ws", s.listen(driven.UI, driver.Server))

	log.Debug().Str("port", s.port).Msg("web socket server listening...")
	go http.ListenAndServe(s.port, nil)
}
