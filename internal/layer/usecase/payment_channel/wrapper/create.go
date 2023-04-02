package payment_channel_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/payment_channel/request"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Create(ctx context.Context, req *request.Create) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Create(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("unable to create")
		trace.RecordError(span, err)
	}

	return err
}
