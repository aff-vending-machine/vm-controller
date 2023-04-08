package idle

import "github.com/aff-vending-machine/vm-controller/internal/core/flow"

func (s *stageImpl) OnInit(c *flow.Ctx) {
	c.Reset()

	machine, _ := s.machineRepo.FindOne(c.UserCtx, []string{"id:=:1"})
	c.Machine = machine

	s.bg(c)
	s.show(c)
}
