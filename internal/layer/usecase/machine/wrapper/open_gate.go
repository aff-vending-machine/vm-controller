package machine_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) OpenGate(ctx context.Context) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.OpenGate(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to open machine gate")
		trace.RecordError(span, err)
	}

	return err
}
