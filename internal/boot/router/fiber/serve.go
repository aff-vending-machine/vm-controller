package fiber

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(driver registry.HTTPTransport) {
	v1 := s.App.Group("/api/v1")
	routeSlot(v1, driver.Slot)

	go s.App.Listen(s.Address)

	log.Info().Str("address", s.Address).Msg("http server listen")
}
