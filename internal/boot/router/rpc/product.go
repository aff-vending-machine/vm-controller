package rpc

import (
	"vm-controller/internal/layer/transport/rpc"
)

func (s *server) routeProduct(endpoint rpc.Product) {
	s.Register("product.set", endpoint.Set)
}
