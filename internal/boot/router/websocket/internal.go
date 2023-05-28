package websocket

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	wsc "vm-controller/internal/layer/service/websocket"
	wss "vm-controller/internal/layer/transport/websocket"

	"github.com/rs/zerolog/log"
)

func (s *serverImpl) listen(client wsc.Frontend, server wss.Frontend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()
		conn, err := s.App.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Err(err).Msg("failed to client upgrade")
			return
		}
		defer conn.Close()

		client.SetConnection(ctx, conn)

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
