package registry

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/transport/http"
	slot_http "github.com/aff-vending-machine/vm-controller/internal/layer/transport/http/slot"
)

// Interface Adapter layers (driver)
type AppDriver struct {
	HTTP AppHTTPDriver
}

type AppHTTPDriver struct {
	Slot http.Slot
}

func NewAppDriver(uc AppUsecase, fw AppFlow) AppDriver {
	return AppDriver{
		AppHTTPDriver{
			slot_http.New(uc.Slot),
		},
	}
}
