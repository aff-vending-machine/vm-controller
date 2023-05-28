package order

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	slots, err := s.slotRepo.FindMany(c.UserCtx, []string{"code:SORT:asc", "is_enable:=:true:bool", "stock:>:0:int"})
	if err != nil {
		s.frontendWs.SendError(c.UserCtx, "order", err.Error())
		log.Error().Err(err).Msg("unable to find all slots")
		return
	}
	s.slots = slots
	s.frontendWs.SendSlots(c.UserCtx, slots)
}
