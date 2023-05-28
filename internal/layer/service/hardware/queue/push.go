package queue

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (hw *hardwareImpl) Push(ctx context.Context, key string, event hardware.Event) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	value := event.ToValueCode()
	log.Debug().Str("key", key).Str("event", event.ToValueCode()).Msg("EVENT: PUSH")
	cmd := hw.client.LPush(ctx, key, value)
	if err := cmd.Err(); err != nil {
		return err
	}
	hw.stacks[event.UID] = &event

	return nil
}
