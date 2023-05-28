package hardware

import (
	"context"

	"vm-controller/internal/core/domain/hardware"
)

type Queue interface {
	Push(ctx context.Context, key string, data hardware.Event) error
	Pop(ctx context.Context, key string) (*hardware.Event, error)
	Clear(ctx context.Context) error

	PushCommand(ctx context.Context, key string, val string) error
	Polling(ctx context.Context, key string, total int, handler hardware.QueueHandler)
	ClearStack(ctx context.Context)
}
