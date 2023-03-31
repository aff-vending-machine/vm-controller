package flow

import "github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/hardware"

func (c *Ctx) AddWaitingEvent(event hardware.Event) {
	c.Events[event.UID] = &event
}

func (c *Ctx) ClearEvent(uid string) {
	c.Events[uid] = nil
	delete(c.Events, uid)
}
