package idle

import "github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"

func (s *Stage) OnInit(c *flow.Ctx) {
	c.Reset()

	machine, _ := s.machineRepo.FindOne(c.UserCtx, []string{"id:=:1"})
	c.Machine = machine

	s.console()
}
