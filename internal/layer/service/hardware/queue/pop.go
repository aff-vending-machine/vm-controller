package queue

import (
	"context"
	"fmt"

	"vm-controller/internal/core/domain/hardware"

	"github.com/rs/zerolog/log"
)

func (h *hardwareImpl) Pop(ctx context.Context, key string) (*hardware.Event, error) {
	cmd := h.client.LPop(ctx, key)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	event := hardware.NewEventFromString(result)
	log.Debug().Str("key", key).Str("event", result).Msg("EVENT: POP")
	if event == nil {
		return nil, fmt.Errorf("invalid event %s", result)
	}

	return event, nil
}
