package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/websocket"
)

// Interface Adapter layers (driver)
type Transport struct {
	RPC       RPCTransport
	WebSocket WebSocketTransport
}

type RPCTransport struct {
	Machine     rpc.Machine
	Slot        rpc.Slot
	Transaction rpc.Transaction
}

type WebSocketTransport struct {
	Frontend websocket.Frontend
}
