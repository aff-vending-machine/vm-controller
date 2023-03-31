package transaction_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Cancel(ctx context.Context, req *request.Cancel) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Cancel(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("unable to cancel order")
		trace.RecordError(span, err)
	}

	return err
}
