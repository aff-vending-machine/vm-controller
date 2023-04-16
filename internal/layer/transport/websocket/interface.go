package websocket

type Frontend interface {
	OnReceived(data []byte) error
}
