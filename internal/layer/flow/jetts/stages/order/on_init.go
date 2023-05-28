package order

import (
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	slots, err := s.slotRepo.FindMany(c.UserCtx, db.NewQuery().AddWhere("is_enable = ?", true).AddWhere("stock > ?", 0).SetOrder("code"))
	if err != nil {
		s.frontendWs.SendError(c.UserCtx, "order", err.Error())
		log.Error().Err(err).Msg("unable to find all slots")
		return
	}
	s.slots = slots
	s.frontendWs.SendSlots(c.UserCtx, slots)
}
