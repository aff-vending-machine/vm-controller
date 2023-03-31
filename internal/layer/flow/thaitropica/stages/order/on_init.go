package order

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *Stage) OnInit(c *flow.Ctx) {
	s.queue.ClearStack(c.UserCtx)

	slots, err := s.slotRepo.FindMany(c.UserCtx, []string{"code:ORDER:asc", "is_enable:=:true", "stock:>:0"})
	if err != nil {
		log.Error().Err(err).Msg("unable to find all slots")
		return
	}

	s.slots = slots

	s.ui.SendSlots(c.UserCtx, slots)
}
