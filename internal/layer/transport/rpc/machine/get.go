package machine

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/module/rabbitmq/rpc"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (r *rpcImpl) Get(c *rpc.Ctx) error {
	ctx, span := trace.Start(c.UserContext)
	defer span.End()

	// usecase execution
	res, err := r.usecase.Get(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to get machine")
		trace.RecordError(span, err)
		return c.InternalServer(err)
	}

	return c.Ok(res)
}
