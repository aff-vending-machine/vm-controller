package lugentpay

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/lugentpay"
	"github.com/rs/zerolog/log"
)

const THAIQR_ENDPOINT = "/cubems/msh/bpms/lugentpay/qrrequest/thaiqr"

func (a *apiImpl) ThaiQR(ctx context.Context, channel *entity.PaymentChannel, data *lugentpay.QRCodeGenerateRequest) (*lugentpay.QRCodeGenerateResponse, error) {
	// body is struct, ignore error
	breq, _ := json.Marshal(data)

	// Set HTTP request
	url := "https://" + channel.Host + THAIQR_ENDPOINT
	log.Info().Str("URL", url).Send()
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(breq))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := lugentpay.QRCodeGenerateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	log.Info().Str("channel", "promptpay").Msg("POST done")

	return &result, nil
}
