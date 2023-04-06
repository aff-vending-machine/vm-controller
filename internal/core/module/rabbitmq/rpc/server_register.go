package rpc

func (c *Server) Register(routingKey string, handler Handler) {
	c.stacks[routingKey] = &handler
}
