package slot_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/response"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) GetOne(ctx context.Context, id uint) (*response.Slot, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	res, err := w.usecase.GetOne(ctx, id)
	if err != nil {
		log.Error().Uint("id", id).Err(err).Msg("unable to get slot")
		trace.RecordError(span, err)
	}

	return res, err
}
