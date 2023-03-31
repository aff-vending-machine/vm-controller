package order

import (
	"encoding/json"
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
)

type WSReceived struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type item struct {
	SlotCode string `json:"code"`
	Quantity int    `json:"quantity"`
}

func (s *Stage) OnWSReceived(c *flow.Ctx, b []byte) error {
	var req WSReceived
	err := json.Unmarshal(b, &req)
	if err != nil {
		return errors.Wrap(err, "convert to struct failed")
	}

	switch req.Action {
	case "refresh":
		slots, err := s.slotRepo.FindMany(c.UserCtx, []string{"code:ORDER:asc", "is_enable:=:true", "stock:>:0"})
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "failed to find all slots")
		}
		s.slots = slots
		s.ui.SendSlots(c.UserCtx, slots)
		s.console(c, req.Action)

		return nil

	case "add":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionAddItem(c, item)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "add item to cart failed")
		}

	case "remove":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionRemoveItem(c, item)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "remove item from cart failed")
		}

	case "set":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionSetItem(c, item)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "set item to cart failed")
		}

	case "clear":
		var item item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &item)
		err := s.actionClearItem(c, item)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "clear item from cart failed")
		}

	case "set-cart":
		var cart []item
		b, _ := json.Marshal(req.Data)
		json.Unmarshal(b, &cart)
		err := s.actionSetCart(c, cart)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "set cart failed")
		}

	case "clear-cart":
		err := s.actionClearCart(c)
		if err != nil {
			s.ui.SendError(c.UserCtx, "order", err.Error())
			s.console(c, req.Action)
			return errors.Wrap(err, "clear cart failed")
		}

	case "done":
		c.ChangeStage <- "payment_channel"

	case "wakeup":
		c.Reset()
		c.ChangeStage <- "order"
		return nil

	default:
		s.ui.SendError(c.UserCtx, "order", fmt.Sprintf("invalid action %s", req.Action))
		return nil
	}

	s.ui.SendCart(c.UserCtx, c.Data.Cart)
	s.console(c, req.Action)

	return nil
}
