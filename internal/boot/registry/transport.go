package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/http"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/keypad"
)

// Interface Adapter layers (driver)
type Transport struct {
	HTTP   HTTPTransport
	Keypad KeypadTransport
}

type HTTPTransport struct {
	Slot http.Slot
}

type KeypadTransport struct {
	Keypad keypad.InputKey
}
