package rpc

import (
	"vm-controller/internal/layer/transport/rpc"
)

func (s *server) routeMachine(endpoint rpc.Machine) {
	s.Register("machine.get", endpoint.Get)
}
