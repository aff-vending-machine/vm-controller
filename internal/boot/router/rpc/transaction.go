package rpc

import (
	"vm-controller/internal/layer/transport/rpc"
)

func (s *server) routeTransaction(endpoint rpc.Transaction) {
	s.Register("transaction.get", endpoint.Get)
	s.Register("transaction.clear", endpoint.Clear)
}
