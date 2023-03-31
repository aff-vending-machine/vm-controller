package emergency

import "github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"

func (s *Stage) OnInit(c *flow.Ctx) {
	c.Reset()
	// debug.Emergency(c)
}
