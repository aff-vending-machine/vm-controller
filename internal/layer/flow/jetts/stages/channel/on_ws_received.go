package channel

import (
	"encoding/json"
	"fmt"

	"vm-controller/internal/core/flow"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type WSReceived struct {
	Action string             `json:"action"`
	Data   PaymentChannelData `json:"data"`
}

type PaymentChannelData struct {
	PaymentChannel string `json:"payment_channel"`
}

func (s *stageImpl) OnWSReceived(c *flow.Ctx, b []byte) error {
	var req WSReceived
	err := json.Unmarshal(b, &req)
	if err != nil {
		return errors.Wrap(err, "convert to struct failed")
	}

	switch req.Action {
	case "confirm":
		for _, channel := range s.channels {
			if req.Data.PaymentChannel == channel.Channel {
				c.PaymentChannel = &channel
				s.createTransaction(c)
				c.ChangeStage <- flow.PAYMENT_STAGE
				return nil
			}
		}

		log.Error().Str("channel", req.Data.PaymentChannel).Msg("unable to select payment channel")
		s.frontendWs.SendError(c.UserCtx, flow.CHANNEL_STAGE, fmt.Sprintf("unable to select payment channel %s", req.Data.PaymentChannel))
		return fmt.Errorf("unable to select payment channel")

	case "wakeup":
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	case "cancel":
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	default:
		log.Error().Str("action", req.Action).Msg("invalid action")
		s.frontendWs.SendError(c.UserCtx, flow.CHANNEL_STAGE, fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}

}
