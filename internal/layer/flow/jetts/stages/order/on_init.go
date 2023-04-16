package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	slots, err := s.slotRepo.FindMany(c.UserCtx, []string{"code:ORDER:asc", "is_enable:=:true", "stock:>:0"})
	if err != nil {
		log.Error().Err(err).Msg("unable to find all slots")
		return
	}
	s.slots = slots
	s.frontendWs.SendSlots(c.UserCtx, slots)

	s.bg(c)
	s.show(c)
}
