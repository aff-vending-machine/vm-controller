package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/http"
	slot_http "github.com/aff-vending-machine/vm-controller/internal/layer/transport/http/slot"
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/ws"
	server_ws "github.com/aff-vending-machine/vm-controller/internal/layer/transport/ws/server"
)

// Interface Adapter layers (driver)
type AppDriver struct {
	HTTP      AppHTTPDriver
	WebSocket WebSocketDriver
}

type AppHTTPDriver struct {
	Slot http.Slot
}

type WebSocketDriver struct {
	Server ws.Server
}

func NewAppDriver(uc AppUsecase, fw AppFlow) AppDriver {
	return AppDriver{
		AppHTTPDriver{
			slot_http.New(uc.Slot),
		},
		WebSocketDriver{
			server_ws.New(fw.ThaiTropica),
		},
	}
}
