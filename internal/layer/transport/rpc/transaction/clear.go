package transaction

import (
	"encoding/json"

	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/transaction/request"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Clear(c *rabbitmq.Ctx) error {
	ctx, span := trace.Start(c.UserContext)
	defer span.End()

	req, err := makeClearRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		trace.RecordError(span, err)
		return c.BadRequest(err)
	}

	// usecase execution
	err = r.usecase.Clear(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to clear transactions")
		trace.RecordError(span, err)
		return c.InternalServer(err)
	}

	return c.Ok(nil)
}

func makeClearRequest(c *rabbitmq.Ctx) (*request.Clear, error) {
	var req request.Clear
	err := json.Unmarshal(c.Delivery.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
