package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc"
)

func (s *server) routeSlot(endpoint rpc.Slot) {
	s.Register("slot.get", endpoint.Get)
	s.Register("slot.set", endpoint.Set)
}
