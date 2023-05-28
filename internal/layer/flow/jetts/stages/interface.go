package stages

import (
	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"
)

type Stage interface {
	OnInit(*flow.Ctx)
	OnEvent(*flow.Ctx, *hardware.Event) error
	OnWSReceived(*flow.Ctx, []byte) error
}
