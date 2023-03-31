package payment_channel_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/payment_channel/response"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Get(ctx context.Context, filter []string) (*response.PaymentChannel, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	res, err := w.usecase.Get(ctx, filter)
	if err != nil {
		log.Error().Strs("filters", filter).Err(err).Msg("unable to get slots")
		trace.RecordError(span, err)
	}

	return res, err
}
