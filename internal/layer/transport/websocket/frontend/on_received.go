package frontend

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (ws *wsImpl) OnReceived(data []byte) error {
	ctx, span := trace.Start(context.Background())
	defer span.End()

	// usecase execution
	err := ws.usecase.OnWSReceived(ctx, data)
	if err != nil {
		log.Error().Err(err).Msg("unable to receive data from web socket")
		trace.RecordError(span, err)
		return err
	}

	return err
}
