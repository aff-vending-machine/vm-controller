package machine_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (uc *wrapperImpl) StartUp(ctx context.Context, req *request.StartUp) (string, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	sn, err := uc.usecase.StartUp(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("unable to startup")
		trace.RecordError(span, err)
	}

	return sn, err
}
