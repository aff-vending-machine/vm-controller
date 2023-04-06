package rpc

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc"
)

func (s *server) routeMachine(endpoint rpc.Machine) {
	s.Register("machine", endpoint.Get)
}
