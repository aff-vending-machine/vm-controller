package emergency

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) bg(c *flow.Ctx) {
	s.displayUc.Background(c.UserCtx, "emergency")
}

func (s *stageImpl) show(c *flow.Ctx) {
	log.Info().Str("stage", "emergency").Msg("LOCKED: press * 7 times to unlock")
}
