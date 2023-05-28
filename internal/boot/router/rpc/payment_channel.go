package rpc

import (
	"vm-controller/internal/layer/transport/rpc"
)

func (s *server) routePaymentChannel(endpoint rpc.PaymentChannel) {
	s.Register("channel.get", endpoint.Get)
	s.Register("channel.set", endpoint.Set)
}
