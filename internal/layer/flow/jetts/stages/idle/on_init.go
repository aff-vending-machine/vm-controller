package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	c.Reset()

	machine, err := s.machineRepo.FindOne(c.UserCtx, db.NewQuery().AddWhere("id = ?", 1))
	if err != nil {
		log.Error().Err(err).Msg("unable to find machine")
	}
	c.Machine = machine
}
