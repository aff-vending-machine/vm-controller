package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/keypad"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/rpc"
)

// Interface Adapter layers (driver)
type Transport struct {
	Keypad KeypadTransport
	RPC    RPCTransport
}

type KeypadTransport struct {
	Keypad keypad.InputKey
}

type RPCTransport struct {
	Machine rpc.Machine
	Slot    rpc.Slot
}
