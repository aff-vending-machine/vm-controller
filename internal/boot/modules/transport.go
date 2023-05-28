package modules

import (
	"vm-controller/internal/layer/transport/rpc"
	"vm-controller/internal/layer/transport/websocket"
)

// Interface Adapter layers (driver)
type Transport struct {
	RPC       RPCTransport
	WebSocket WebSocketTransport
}

type RPCTransport struct {
	Machine        rpc.Machine
	PaymentChannel rpc.PaymentChannel
	Product        rpc.Product
	Slot           rpc.Slot
	Transaction    rpc.Transaction
}

type WebSocketTransport struct {
	Frontend websocket.Frontend
}
