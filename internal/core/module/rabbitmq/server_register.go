package rabbitmq

func (s *Server) Register(routingKey string, handler Handler) {
	s.stacks[routingKey] = &handler
}
