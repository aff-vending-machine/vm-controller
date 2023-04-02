package machine_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Emergency(ctx context.Context) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Emergency(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to set emergency stage")
		trace.RecordError(span, err)
	}

	return err
}
