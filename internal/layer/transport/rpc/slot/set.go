package slot

import (
	"encoding/json"

	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq/rpc"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Set(c *rpc.Ctx) error {
	ctx, span := trace.Start(c.UserContext)
	defer span.End()

	req, err := makeSetRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		trace.RecordError(span, err)
		return c.BadRequest(err)
	}

	// usecase execution
	err = r.usecase.Set(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to set all slots")
		trace.RecordError(span, err)
		return c.InternalServer(err)
	}

	return c.Ok(nil)
}

func makeSetRequest(c *rpc.Ctx) (*request.Set, error) {
	var req request.Set
	err := json.Unmarshal(c.Delivery.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
