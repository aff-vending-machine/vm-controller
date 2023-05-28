package queue

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (h *hardwareImpl) PushCommand(ctx context.Context, key string, data string) error {
	log.Debug().Str("key", key).Str("data", data).Msg("HWCMD: PUSH")
	cmd := h.client.LPush(ctx, key, data)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}
