package queue_hardware

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (h *hardwareImpl) Polling(ctx context.Context, key string, total int, handle hardware.QueueHandler) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	count := 0
	for count < total {
		if len(h.stacks) == 0 {
			break
		}

		event, err := h.Pop(ctx, key)
		if err != nil {
			continue
		}

		if h.stacks[event.UID] == nil {
			log.Error().Str("key", key).Str("uid", event.UID).Msg("pop event is nil")
			continue
		}

		h.stacks[event.UID] = event
		if err := handle(event); err != nil {
			log.Error().Str("event", event.ToValueCode()).Err(err).Msg("handle event is error")
			continue
		}

		count++
		time.Sleep(100 * time.Millisecond)
	}

	h.Clear(ctx)
	h.stacks = make(map[string]*hardware.Event)
}
