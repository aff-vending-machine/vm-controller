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
	case "cancel", "done", "new-order":
		if s.polling {
			log.Warn().Str("order_id", c.Data.MerchantOrderID).Str("action", req.Action).Str("stage", "receive").Msg("Cancel items")
			s.status = CANCEL
			s.queue.ClearStack(c.UserCtx)
			s.polling = false
		} else {
			c.Reset()
			c.ChangeStage <- "idle"
		}
		return nil

	case "open-gate":
		s.status = 0
		s.queue.PushCommand(c.UserCtx, "COMMAND", "OPEN_GATE")
		return nil

	case "wakeup":
		if s.polling {
			log.Warn().Str("order_id", c.Data.MerchantOrderID).Str("action", req.Action).Str("stage", "receive").Msg("Please grab item")
			s.ui.SendGrabItem(c.UserCtx, "receive", "Please grab item")
		} else {
			c.Reset()
			c.ChangeStage <- "order"
		}
		return nil

	default:
		log.Warn().Str("action", req.Action).Str("stage", "receive").Msg("invalid action")
		// s.ui.SendError(c.UserCtx, "receive", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}
}
