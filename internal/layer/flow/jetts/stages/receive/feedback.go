package receive

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
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
			return s.errorFeedback(c, event)
		}

		s.status = WAIT
		codeFilter := []string{fmt.Sprintf("code:=:%s", event.SlotCode)}
		slot, err := s.slotRepo.FindOne(c.UserCtx, codeFilter)
		if err != nil {
			log.Error().Strs("code filters", codeFilter).Err(err).Msg("unable to find slot")
			return err
		}

		slot.Stock -= 1
		data := map[string]interface{}{
			"stock": slot.Stock,
		}
		_, err = s.slotRepo.UpdateMany(c.UserCtx, codeFilter, data)
		if err != nil {
			log.Error().Strs("code filters", codeFilter).Interface("data", data).Err(err).Msg("unable to update slot")
			return err
		}

		for i, item := range c.Data.Cart {
			if item.SlotCode == event.SlotCode {
				c.Data.Cart[i].Received += 1
				log.Info().Str("slot", item.SlotCode).Msg("update slot")
				break
			}
		}

		c.ClearEvent(event.UID)
		s.show(c)
		return nil
	}
}

func (s *stageImpl) errorFeedback(c *flow.Ctx, event *hardware.Event) error {
	switch event.Status {
	case "E0":
		log.Warn().Str("event", event.ToValueCode()).Str("slot_code", event.SlotCode).Msg("Item is not drop, maybe this slot has no item")
		// no item
		s.queue.Clear(c.UserCtx)
		s.queue.PushCommand(c.UserCtx, "COMMAND", "RESET")
		s.slotRepo.UpdateMany(c.UserCtx, []string{fmt.Sprintf("code:=:%s", event.SlotCode)}, map[string]interface{}{"is_enable": false})

		err := fmt.Errorf("hardware error: item is not drop, maybe slot %s has no item (event: %s)", event.SlotCode, event.ToValueCode())
		s.error(c, err, "this slot has no item")
		s.updateErrorTransaction(c, err)
		s.status = E0
		return flow.ErrMachineE0

	case "E1":
		// don't get item
		log.Warn().Str("event", event.ToValueCode()).Msg("Customer don't grab item")

		err := fmt.Errorf("hardware error: customer don't grab item (event: %s)", event.ToValueCode())
		s.error(c, err, "please grab item")
		s.updateErrorTransaction(c, err)
		s.status = E1
		return flow.ErrMachineE1

	case "E2":
		// unknown
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: CRITICAL ERROR (event: %s)", event.ToValueCode()))
		s.status = E2
		return flow.ErrMachineE2

	default:
		s.updateErrorTransaction(c, fmt.Errorf("hardware error: CRITICAL ERROR (event: %s)", event.ToValueCode()))
		s.status = E2
		return flow.ErrMachineE2
	}
}
