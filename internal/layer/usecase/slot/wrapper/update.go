package slot_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Update(ctx context.Context, req *request.Update) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Update(ctx, req)
	if err != nil {
		log.Error().
			Interface("request", req).
			Err(err).
			Msg("unable to update slot")
		trace.RecordError(span, err)
	}

	return err
}
