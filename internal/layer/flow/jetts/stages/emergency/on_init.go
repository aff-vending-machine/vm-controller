package emergency

import "vm-controller/internal/core/flow"

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.reset = 0
	c.Reset()
}
