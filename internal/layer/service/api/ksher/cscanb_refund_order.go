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

func (c *apiImpl) RefundOrder(ctx context.Context, channel *entity.PaymentChannel, orderID string, body *ksher.RefundOrderBody) (*ksher.RefundOrderResult, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	path := utils.GenerateURLPath(CSCANB_PATH, orderID) // no prefix "/" after gen
	pregen := toJson(body)
	signature := generateSignature(path, pregen, channel.Token)
	body.Signature = signature

	// body is struct, ignore error
	breq, _ := json.Marshal(body)

	// Set HTTP request
	url := utils.GenerateURLPath(channel.Host, CSCANB_PATH, orderID)
	log.Debug().Str("channel", "ksher").Str("URL", url).Str("signature", signature).Msg("PUT refund order")

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(breq))
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

	// v is not nil and be a pointer, ignore error
	result := ksher.RefundOrderResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
