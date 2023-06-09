package receive

import (
	"fmt"
	"time"

	"vm-controller/internal/core/domain/hardware"
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) feedback(c *flow.Ctx) hardware.QueueHandler {
	return func(event *hardware.Event) error {
		log.Info().Str("event", event.ToValueCode()).Msg("Feedback")
		if c.Events[event.UID] == nil {
			c.ClearEvent(event.UID)
			log.Warn().Str("uid", event.UID).Msg("unable to find event handler")
			return flow.ErrInvalidEvent
		}

		if event.Status != "S0" {
			log.Warn().Str("event", event.ToValueCode()).Msg("hardware error")
			err := s.errorFeedback(c, event)
			return err
		}

		s.status = WAIT
		codeFilter := db.NewQuery().AddWhere("code = ?", event.SlotCode)
		slot, err := s.slotRepo.FindOne(c.UserCtx, codeFilter)
		if err != nil {
			log.Error().Interface("filter", codeFilter).Err(err).Msg("unable to find slot")
			return err
		}

		slot.Stock -= 1
		data := map[string]interface{}{
			"stock": slot.Stock,
		}
		_, err = s.slotRepo.Update(c.UserCtx, codeFilter, data)
		if err != nil {
			log.Error().Interface("filter", codeFilter).Interface("data", data).Err(err).Msg("unable to update slot")
			return err
		}

		for i, item := range c.Data.Cart {
			if item.SlotCode == event.SlotCode {
				c.Data.Cart[i].Received += 1
				log.Info().Str("slot", item.SlotCode).Msg("update slot")
				s.frontendWs.SendReceivedItem(c.UserCtx, c.Data.MerchantOrderID, c.Data.Cart, item)
				break
			}
		}

		c.ClearEvent(event.UID)
		log.Info().Str("stage", string(c.Stage)).Int("remaining", len(c.Events)).Str("order_id", c.Data.MerchantOrderID).Interface("events", c.Events).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Float64("Price", c.Data.TotalPrice()).Float64("Pay", c.Data.TotalPay()).Msg("SLOG: receive event")
		return nil
	}
}

func (s *stageImpl) errorFeedback(c *flow.Ctx, event *hardware.Event) error {
	switch event.Status {
	case "E0":
		// no item
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: item is not drop, maybe slot %s has no item (event: %s)", event.SlotCode, event.ToValueCode()))
		s.queue.ClearStack(c.UserCtx)
		s.queue.PushCommand(c.UserCtx, "COMMAND", "RESET")
		log.Warn().Str("event", event.ToValueCode()).Str("slot_code", event.SlotCode).Msg("Item is not drop, maybe this slot has no item")
		s.slotRepo.Update(c.UserCtx, db.NewQuery().AddWhere("code = ?", event.SlotCode), map[string]interface{}{"is_enable": false})
		time.Sleep(5 * time.Second)
		s.queue.PushCommand(c.UserCtx, "COMMAND", "RESET")
		s.frontendWs.SendError(c.UserCtx, c.Stage, "Please Contact Center")
		return flow.ErrMachineE0

	case "E1":
		// don't get item
		s.status = 0
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: customer don't grab item (event: %s)", event.ToValueCode()))
		log.Warn().Str("event", event.ToValueCode()).Msg("Customer don't grab item, auto open gate")
		s.queue.PushCommand(c.UserCtx, "COMMAND", "OPEN_GATE")
		return flow.ErrMachineE1

	case "E2":
		// unknown
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: CRITICAL ERROR (event: %s)", event.ToValueCode()))
		s.frontendWs.SendEmergency(c.UserCtx, flow.ErrMachineE2)
		return flow.ErrMachineE2

	default:
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: CRITICAL ERROR (event: %s)", event.ToValueCode()))
		s.frontendWs.SendEmergency(c.UserCtx, flow.ErrMachineE2)
		return flow.ErrMachineE2
	}
}
