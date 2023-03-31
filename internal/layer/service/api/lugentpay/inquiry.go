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

const INQUIRY_ENDPOINT = "/cubems/msh/rest/lugentpay/checkpayment"

func (a *apiImpl) Inquiry(ctx context.Context, channel *entity.PaymentChannel, data *lugentpay.InquiryBody) (*lugentpay.InquiryResult, error) {
	// body is struct, ignore error
	breq, _ := json.Marshal(data)

	// Set HTTP request
	url := "https://" + channel.Host + INQUIRY_ENDPOINT
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

	result := lugentpay.InquiryResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
