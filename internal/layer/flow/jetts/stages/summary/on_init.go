package summary

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	ts := time.Now().Format("20060102150405")
	// prevent duplicate other machine
	c.Data.MerchantOrderID = fmt.Sprintf("%s%s", c.Machine.Name, ts)

	s.bg(c)
	s.show(c)
}
