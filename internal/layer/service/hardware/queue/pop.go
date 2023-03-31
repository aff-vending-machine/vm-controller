package queue_hardware

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (h *hardwareImpl) Pop(ctx context.Context, key string) (*hardware.Event, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	cmd := h.client.LPop(ctx, key)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	event := hardware.NewEventFromString(result)
	log.Debug().Str("key", key).Str("event", result).Msg("pop event")
	if event == nil {
		return nil, fmt.Errorf("invalid event %s", result)
	}

	return event, nil
}
