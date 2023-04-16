package payment

import (
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
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
		s.updatePaidTransaction(c)

		s.frontendWs.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
		c.ChangeStage <- "receive"
		return nil

	case "cancel":
		if s.CancelFn != nil {
			s.CancelFn()
			s.CancelFn = nil
		}

		s.updateCancelTransaction(c, "user")
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "wakeup":
		s.updateCancelTransaction(c, "machine")
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	default:
		s.frontendWs.SendError(c.UserCtx, "payment", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
