package stages

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

type Stage interface {
	OnInit(*flow.Ctx)
	OnEvent(*flow.Ctx, *hardware.Event) error
	OnKeyPressed(*flow.Ctx, hardware.Key) error
}
