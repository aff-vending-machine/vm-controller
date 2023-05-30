package payment

import (
	"encoding/json"
	"fmt"

	"vm-controller/internal/core/flow"

	"github.com/pkg/errors"
)

type WSReceived struct {
	Action string `json:"action"`
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

	case "cancel":
		if s.CancelFn != nil {
			s.CancelFn()
			s.CancelFn = nil
		}

		s.updateCancelTransaction(c, "user")
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	case "wakeup":
		s.updateCancelTransaction(c, "machine")
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	case "bypass":
		// s.bypass = true
		return nil

	default:
		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
