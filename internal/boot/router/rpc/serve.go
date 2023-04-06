package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(queue string, driver registry.RPCTransport) {
	s.routeMachine(driver.Machine)
	s.routeSlot(driver.Slot)

	go s.Listen(queue)

	log.Info().Str("queue", queue).Msg("rpc server listen")
}
