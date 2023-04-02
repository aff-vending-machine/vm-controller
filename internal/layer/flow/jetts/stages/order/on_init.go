package order

import (
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.queue.ClearStack(c.UserCtx)

	s.bg(c)
	s.show(c)
}
