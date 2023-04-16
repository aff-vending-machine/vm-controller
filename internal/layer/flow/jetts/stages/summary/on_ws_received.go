package summary

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
	case "cancel":
		c.ChangeStage <- "order"
		return nil

	case "confirm":
		err := s.createTransaction(c)
		if err != nil {
			c.ChangeStage <- "emergency"
			return s.error(c, flow.ErrOutOfService, "out of service")
		}

		c.ChangeStage <- "payment_channel"
		return nil

	default:
		s.frontendWs.SendError(c.UserCtx, "payment", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
