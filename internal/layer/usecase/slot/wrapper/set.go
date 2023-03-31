package slot_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Set(ctx context.Context, s *request.Set) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Set(ctx, s)
	if err != nil {
		log.Error().
			Interface("request", s).
			Err(err).
			Msg("unable to set a slot")
		trace.RecordError(span, err)
	}

	return err
}
