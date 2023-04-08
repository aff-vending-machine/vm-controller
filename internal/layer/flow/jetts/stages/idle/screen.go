package idle

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "idle")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().Str("stage", "idle").Msg("idle stage")
	s.displayUc.StageIdle(c.UserCtx)
}
