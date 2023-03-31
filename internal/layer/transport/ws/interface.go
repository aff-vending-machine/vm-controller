package ws

type Server interface {
	OnReceived(data []byte) error
}
