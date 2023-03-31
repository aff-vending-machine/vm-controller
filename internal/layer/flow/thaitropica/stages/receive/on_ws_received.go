package receive

import (
	"encoding/json"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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
	case "cancel":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "done":
		c.Reset()
		c.ChangeStage <- "idle"
		return nil

	case "new-order":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	case "open-gate":
		s.queue.PushCommand(c.UserCtx, "COMMAND", "OPEN_GATE")
		s.reset()
		return nil

	case "wakeup":
		if !s.polling {
			c.Reset()
			c.ChangeStage <- "order"
		} else {
			log.Warn().Str("action", req.Action).Str("stage", "receive").Msg("Please grab item")
			s.ui.SendGrabItem(c.UserCtx, "receive", "Please grab item")
		}
		return nil

	default:
		log.Warn().Str("action", req.Action).Str("stage", "receive").Msg("invalid action")
		// s.ui.SendError(c.UserCtx, "receive", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
