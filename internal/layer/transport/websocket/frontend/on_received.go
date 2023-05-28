package frontend

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (ws *wsImpl) OnReceived(data []byte) error {
	ctx := context.Background()

	// usecase execution
	err := ws.usecase.OnWSReceived(ctx, data)
	if err != nil {
		log.Error().Err(err).Msg("unable to receive data from web socket")
		return err
	}

	return err
}
