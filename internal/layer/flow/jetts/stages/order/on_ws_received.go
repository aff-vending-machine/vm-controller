package order

import (
	"encoding/json"
	"fmt"

	"vm-controller/internal/core/flow"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type WSReceived struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type item struct {
	SlotCode string `json:"code"`
	Quantity int    `json:"quantity"`
}

func (s *stageImpl) OnWSReceived(c *flow.Ctx, b []byte) error {
	var req WSReceived
	err := json.Unmarshal(b, &req)
	if err != nil {
		log.Error().Err(err).Msg("unable to unmarshal request")
		return errors.Wrap(err, "convert to struct failed")
	}

	switch req.Action {
	case "refresh":
		s.OnInit(c)
		return nil

	case "add":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionAddItem(c, item)
		if err != nil {
			log.Error().Err(err).Interface("item", item).Msg("unable to add item")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "add item to cart failed")
		}

	case "remove":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionRemoveItem(c, item)
		if err != nil {
			log.Error().Err(err).Interface("item", item).Msg("unable to remove item")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "remove item from cart failed")
		}

	case "set":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionSetItem(c, item)
		if err != nil {
			log.Error().Err(err).Interface("item", item).Msg("unable to set item")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "set item to cart failed")
		}

	case "clear":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionClearItem(c, item)
		if err != nil {
			log.Error().Err(err).Interface("item", item).Msg("unable to clear item")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "clear item from cart failed")
		}

	case "set-cart":
		var cart []item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &cart)
		err := s.actionSetCart(c, cart)
		if err != nil {
			log.Error().Err(err).Interface("cart", cart).Msg("unable to set cart")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "set cart failed")
		}

	case "clear-cart":
		err := s.actionClearCart(c)
		if err != nil {
			log.Error().Err(err).Msg("unable to clear cart")
			s.frontendWs.SendError(c.UserCtx, c.Stage, err.Error())
			return errors.Wrap(err, "clear cart failed")
		}

	case "done":
		if len(c.Data.Cart) == 0 {
			log.Error().Msg("empty cart")
			s.frontendWs.SendError(c.UserCtx, c.Stage, "empty cart")
			return nil
		}

		c.ChangeStage <- flow.CHANNEL_STAGE
		return nil

	case "wakeup":
		c.Reset()
		c.ChangeStage <- flow.ORDER_STAGE
		return nil

	default:
		s.frontendWs.SendError(c.UserCtx, c.Stage, fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}

	s.frontendWs.SendCart(c.UserCtx, c.Data.Cart)

	return nil
}
