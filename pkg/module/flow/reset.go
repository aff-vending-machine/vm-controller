package flow

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"

func (c *Ctx) Reset() {
	// c.Stage = "idle"
	c.Data.MerchantOrderID = ""
	c.Data.Cart = make([]hardware.Item, 0)
	c.Events = make(map[string]*hardware.Event)
}

func (c *Ctx) Emergency(err error) {
	c.Error = err
	c.ChangeStage <- "error"
}
