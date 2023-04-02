package fiber

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/app/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(port int, driver registry.AppHTTPDriver) {
	v1 := s.app.Group("/api/v1")
	routeSlot(v1, driver.Slot)

	go s.app.Listen(fmt.Sprintf(":%d", port))

	log.Info().Int("port", port).Msg("http server listen")
}
