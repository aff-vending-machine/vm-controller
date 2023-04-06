package flow

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

type Jetts interface {
	ListenEvent(sn string)
	OnInit(ctx context.Context)
	OnEvent(ctx context.Context, event *hardware.Event) error
	OnKeyPressed(ctx context.Context, key string) error
}
