package stages

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
)

type Stage interface {
	OnInit(*flow.Ctx)
	OnEvent(*flow.Ctx, *hardware.Event) error
	OnWSReceived(*flow.Ctx, []byte) error
}
