package board

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (r *raspiImpl) OnPressed(key string) error {
	ctx, span := trace.Start(context.Background())
	defer span.End()

	// execute usecase
	err := r.usecase.OnKeyPressed(ctx, key)
	if err != nil {
		log.Error().Err(err).Msg("unable to execute key pressed event")
		trace.RecordError(span, err)
		return err
	}

	return err
}
