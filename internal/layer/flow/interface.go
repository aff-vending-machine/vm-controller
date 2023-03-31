package flow

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
)

type ThaiTropica interface {
	ListenEvent(ctx context.Context, sn string)
	OnInit(ctx context.Context)
	OnEvent(ctx context.Context, event *hardware.Event) error
	OnWSReceived(ctx context.Context, data []byte) error
}
