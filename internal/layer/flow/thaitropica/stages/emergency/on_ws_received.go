package emergency

import (
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
)

type WSReceived struct {
	Action string `json:"action"`
}

func (s *Stage) OnWSReceived(c *flow.Ctx, b []byte) error {
	var req WSReceived
	err := json.Unmarshal(b, &req)
	if err != nil {
		return errors.Wrap(err, "convert to struct failed")
	}

	switch req.Action {
	case "reset":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "wakeup":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	default:
		s.ui.SendError(c.UserCtx, "emergency", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
