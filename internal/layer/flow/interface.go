package flow

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
)

type Jetts interface {
	ListenEvent(string)
	OnInit(context.Context)
	OnEvent(context.Context, *hardware.Event) error
	OnWSReceived(context.Context, []byte) error
}
