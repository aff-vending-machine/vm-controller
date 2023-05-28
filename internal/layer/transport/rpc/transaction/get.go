package transaction

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/rabbitmq"
	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Get(c *rabbitmq.Ctx) error {
	ctx := c.UserContext

	// usecase execution
	res, err := r.usecase.Get(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to get all transactions")
		return c.InternalServer(err)
	}

	return c.Ok(res)
}
