package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/keypad"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/websocket"
)

// Interface Adapter layers (driver)
type Transport struct {
	Keypad    KeypadTransport
	RPC       RPCTransport
	WebSocket WebSocketTransport
}

type KeypadTransport struct {
	Keypad keypad.InputKey
}

type RPCTransport struct {
	Machine     rpc.Machine
	Slot        rpc.Slot
	Transaction rpc.Transaction
}

type WebSocketTransport struct {
	Frontend websocket.Frontend
}
