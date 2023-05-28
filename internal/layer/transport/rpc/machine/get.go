package machine

import (
	"vm-controller/internal/core/infra/network/rabbitmq"

	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Get(c *rabbitmq.Ctx) error {
	ctx := c.UserContext

	// usecase execution
	res, err := r.usecase.Get(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to get machine")
		return c.InternalServer(err)
	}

	return c.Ok(res)
}
