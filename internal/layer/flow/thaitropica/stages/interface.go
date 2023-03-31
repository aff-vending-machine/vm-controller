package stages

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

type Stage interface {
	OnInit(*flow.Ctx)
	OnEvent(*flow.Ctx, *hardware.Event) error
	OnWSReceived(*flow.Ctx, []byte) error
}
