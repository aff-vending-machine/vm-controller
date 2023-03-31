package smartedcv2

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	"github.com/rs/zerolog/log"
)

func (a *apiImpl) Void(ctx context.Context, salereq *smartedc.VoidRequest) (*smartedc.VoidResult, error) {
	// body is struct, ignore error
	breq, _ := json.Marshal(salereq)

	// Set HTTP request
	log.Info().Str("channel", "creditcard").Interface("request", salereq).Msg("POST request")
	req, err := http.NewRequest(http.MethodPost, "http://rpi-smartedc:8082/api/v1/smartedc/void", bytes.NewBuffer(breq))
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

	result := smartedc.VoidResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	log.Info().Str("channel", "creditcard").Interface("response", result).Msg("POST done")

	return &result, nil
}
