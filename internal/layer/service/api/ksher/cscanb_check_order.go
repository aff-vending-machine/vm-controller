package ksher

import (
	"context"
	"encoding/json"
	"net/http"

	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/domain/ksher"
	"vm-controller/pkg/helpers/gen"

	"github.com/rs/zerolog/log"
)

func (c *apiImpl) CheckOrder(ctx context.Context, channel *entity.PaymentChannel, orderID string, query *ksher.CheckOrderQuery) (*ksher.CheckOrderResult, error) {
	path := "/" + gen.ToURLPath(CSCANB_PATH, orderID) // no prefix "/" after gen
	pregen := toJson(query)
	signature := generateSignature(path, pregen, channel.Token)
	query.Signature = signature

	url := gen.ToURLPath(channel.Host, CSCANB_PATH, orderID)
	log.Debug().Str("channel", "ksher").Str("URL", url).Str("signature", signature).Msg("GET check order")

	// Set HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Cache-Control", "no-cache")

	// Set Queries
	queries := req.URL.Query()

	// query is struct, ignore error
	queryMap := map[string]string{}
	qreq, _ := json.Marshal(query)
	json.Unmarshal(qreq, &queryMap)

	for key, value := range queryMap {
		queries.Add(key, value)
	}

	// assign encoded query string to http request
	req.URL.RawQuery = queries.Encode()

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := ksher.CheckOrderResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
