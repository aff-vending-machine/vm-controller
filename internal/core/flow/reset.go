package flow

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/domain/hardware"
)

func (c *Ctx) Reset() {
	// c.Stage = "idle"
	c.Data.MerchantOrderID = ""
	c.PaymentChannel = &entity.PaymentChannel{}
	c.Data.Cart = make([]hardware.Item, 0)
	c.Events = make(map[string]*hardware.Event)
}

func (c *Ctx) Emergency(err error) {
	if err != nil {
		c.Error = err
		c.ChangeStage <- EMERGENCY_STAGE
	}
}
