package emergency

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
	case "reset":
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	case "wakeup":
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	default:
		s.frontendWs.SendError(c.UserCtx, "emergency", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
