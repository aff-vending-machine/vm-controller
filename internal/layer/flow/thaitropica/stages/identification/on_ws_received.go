package identification

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
	case "request-otp":
		return s.requestOTP(c, req)

	case "resend-otp":
		return s.sendOTP(c, req)

	case "check-otp":
		return s.checkOTP(c, req)

	case "cancel":
		s.updateCancelTransaction(c)
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "wakeup":
		s.updateCancelTransactionByMachine(c)
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	default:
		s.ui.SendError(c.UserCtx, "identification", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
