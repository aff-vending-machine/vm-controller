package ksher

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/aff-vending-machine/vm-controller/pkg/utils"
	"github.com/rs/zerolog/log"
)

func (c *apiImpl) CreateOrder(ctx context.Context, channel *entity.PaymentChannel, body *ksher.CreateOrderBody) (*ksher.CreateOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	pregen := toJson(body)
	// delete(pregen, "provider")
	signature := generateSignature(CSCANB_PATH, pregen, channel.Token)
	body.Signature = signature

	// body is struct, ignore error
	breq, _ := json.Marshal(body)

	// Set HTTP request
	url := utils.GenerateURLPath(channel.Host, CSCANB_PATH)
	log.Debug().Str("channel", "ksher").Str("URL", url).Str("signature", signature).Msg("POST create order")

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(breq))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := ksher.CreateOrderResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
