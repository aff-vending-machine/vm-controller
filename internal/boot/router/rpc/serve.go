package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/boot/modules"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(queue string, driver modules.RPCTransport) {
	s.routeMachine(driver.Machine)
	s.routeSlot(driver.Slot)
	s.routeTransaction(driver.Transaction)

	go s.Listen(queue)

	log.Info().Str("queue", queue).Msg("rpc server listen")
}
