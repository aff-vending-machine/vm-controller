package payment

import (
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
)

type WSReceived struct {
	Action string             `json:"action"`
	Data   IdentificationData `json:"data"`
}

type IdentificationData struct {
	Mail      string `json:"mail,omitempty"`
	Reference string `json:"reference,omitempty"`
	OTP       string `json:"otp,omitempty"`
}

func (s *stageImpl) OnWSReceived(c *flow.Ctx, b []byte) error {
	var req WSReceived
	err := json.Unmarshal(b, &req)
	if err != nil {
		return errors.Wrap(err, "convert to struct failed")
	}

	switch req.Action {
	case "refresh":
		s.OnInit(c)
		return nil

	case "bypass":
		s.onPaid(c)
		return nil

	case "cancel":
		if s.CancelFn != nil {
			s.CancelFn()
			s.CancelFn = nil
		}

		s.onCancel(c)
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "wakeup":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	default:
		s.ui.SendError(c.UserCtx, "payment", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
