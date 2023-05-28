package link2500

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/gen"
	"github.com/rs/zerolog/log"
)

type SettlementResponse struct {
	Code    int                        `json:"code"`
	Status  string                     `json:"status"`
	Data    *link2500.SettlementResult `json:"data,omitempty"`
	Message *string                    `json:"message,omitempty"`
}

func (a *apiImpl) Settlement(ctx context.Context, channel *entity.PaymentChannel, body *link2500.SettlementRequest) (*link2500.SettlementResult, error) {
	// body is struct, ignore error
	breq, _ := json.Marshal(body)

	// Set HTTP request
	url := gen.ToURLPath(channel.Host, LINK2500_PATH, "settlement")

	log.Debug().Str("channel", "creditcard").Str("URL", url).Interface("request", body).Msg("POST settlement")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(breq))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	req = req.WithContext(ctx)
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := SettlementResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	log.Info().Str("channel", "creditcard").Interface("response", result).Int("status", resp.StatusCode).Msg("POST done")

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", result.Code, *result.Message)
	}

	return result.Data, nil
}
