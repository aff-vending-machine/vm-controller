package rabbitmq

type Client struct {
	Conn *Connection
}

func NewClient(conn *Connection) *Client {
	return &Client{
		Conn: conn,
	}
}
