package machine_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Reset(ctx context.Context) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Reset(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to reset machine")
		trace.RecordError(span, err)
	}

	return err
}
