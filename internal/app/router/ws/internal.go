package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	wsc "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/ws"
	wss "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/transport/ws"
	"github.com/rs/zerolog/log"
)

func (s *serverImpl) listen(client wsc.UI, server wss.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()
		conn, err := s.app.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Err(err).Msg("failed to client upgrade")
			return
		}
		defer conn.Close()

		client.SetConnection(ctx, conn)
		// client.SendToIdle(ctx)

		for {
			// Read the response from the client
			var response map[string]interface{}
			err = conn.ReadJSON(&response)
			if err != nil {
				log.Error().Err(err).Msg("failed to client upgrade")
				break
			}
			b, _ := json.Marshal(response)

			// Update the system state based on the client response
			err := server.OnReceived(b)
			if err != nil {
				continue
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}
