package rabbitmq

type Server struct {
	Conn   *Connection
	stacks map[string]*Handler
}

func NewServer(conn *Connection) *Server {
	return &Server{
		Conn:   conn,
		stacks: make(map[string]*Handler),
	}
}
