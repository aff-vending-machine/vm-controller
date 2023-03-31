package slot_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Clear(ctx context.Context, filter []string) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Clear(ctx, filter)
	if err != nil {
		log.Error().
			Strs("filters", filter).
			Err(err).
			Msg("unable to remove slot")
		trace.RecordError(span, err)
	}

	return err
}
