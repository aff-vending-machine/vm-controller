package slot

import (
	"encoding/json"

	"github.com/aff-vending-machine/vm-controller/internal/core/infra/rabbitmq"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Set(c *rabbitmq.Ctx) error {
	ctx := c.UserContext

	req, err := makeSetRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		return c.BadRequest(err)
	}

	// usecase execution
	err = r.usecase.Set(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to set all slots")
		return c.InternalServer(err)
	}

	return c.Ok(nil)
}

func makeSetRequest(c *rabbitmq.Ctx) (*request.Set, error) {
	var req request.Set
	err := json.Unmarshal(c.Delivery.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
