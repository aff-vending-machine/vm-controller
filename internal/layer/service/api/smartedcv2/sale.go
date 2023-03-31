package smartedcv2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	"github.com/rs/zerolog/log"
)

type SaleResponse struct {
	Code    int                  `json:"code"`
	Status  string               `json:"status"`
	Data    *smartedc.SaleResult `json:"data,omitempty"`
	Message *string              `json:"message,omitempty"`
}

func (a *apiImpl) Sale(ctx context.Context, salereq *smartedc.SaleRequest) (*smartedc.SaleResult, error) {
	// body is struct, ignore error
	breq, _ := json.Marshal(salereq)

	// Set HTTP request
	log.Info().Str("channel", "creditcard").Interface("request", salereq).Msg("POST request")
	req, err := http.NewRequest(http.MethodPost, "http://rpi-smartedc:8082/api/v1/smartedc/sale", bytes.NewBuffer(breq))
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

	result := SaleResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", result.Code, *result.Message)
	}

	log.Info().Str("channel", "creditcard").Interface("response", result).Msg("POST done")

	return result.Data, nil
}
