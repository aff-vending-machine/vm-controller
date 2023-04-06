package emergency

import "github.com/aff-vending-machine/vm-controller/internal/core/flow"

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.reset = 0
	c.Reset()

	s.bg(c)
	s.show(c)
}
