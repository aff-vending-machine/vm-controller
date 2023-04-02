package receive

import (
	"fmt"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.queue.Clear(c.UserCtx)

	if err := s.addEvents(c); err != nil {
		log.Error().Err(err).Msg("unable to add events")
		return
	}

	log.Info().Str("stage", "receive").Int("remaining", len(c.Events)).Str("order_id", c.Data.MerchantOrderID).Interface("events", c.Events).Interface("cart", c.Data.Cart).Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Float64("Price", c.Data.TotalPrice()).Float64("Pay", c.Data.TotalPay()).Msg("SLOG: receive event")
	go s.checkEvent(c)
}

func (s *stageImpl) checkEvent(c *flow.Ctx) {
	total := c.Data.TotalQuantity()

	log.Info().Msg("start polling")
	s.polling = true
	s.queue.Polling(c.UserCtx, "RESPONSE", total, s.feedback(c))
	s.polling = false
	log.Info().Msg("stop polling")

	if s.status == CANCEL {
		s.updateCancelTransaction(c)

		s.queue.Clear(c.UserCtx)
		c.ChangeStage <- "idle"
		return
	}

	if c.Data.TotalQuantity() != c.Data.TotalReceived() {
		s.updateBrokenTransaction(c)
		// s.void(c)
	} else {
		s.updateDoneTransaction(c)
		s.customerRepo.UpdateMany(c.UserCtx, []string{fmt.Sprintf("email:=:%s", c.Data.Mail)}, map[string]interface{}{
			"is_received": true,
			"cart":        c.Data.Raw(),
		})
	}

	log.Info().Int("Quantity", c.Data.TotalQuantity()).Int("Received", c.Data.TotalReceived()).Msg("DONE")
	s.ui.SendDone(c.UserCtx, c.Data.MerchantOrderID, c.Data.Cart)

	s.status = DONE
	s.queue.Clear(c.UserCtx)

	c.ChangeStage <- "idle"
}

// func (s *stageImpl) void(c *flow.Ctx) {
// 	switch strings.ToLower(c.PaymentChannel.Channel) {
// 	case "creditcard":
// 		if c.Data.TotalReceived() == 0 {
// 			s.smartedc.Void(c.UserCtx, &smartedc.VoidRequest{
// 				TradeType:        "CARD",
// 				InvoiceNo:        c.Data.InvoiceNo,
// 				CardApprovalCode: c.Data.CardApprovalCode,
// 				Amount:           c.Data.TotalPrice(),
// 				TransactionType:  "VOID",
// 				POSRefNo:         c.Data.MerchantOrderID,
// 			})
// 		}
// 		return
// 	default:
// 	}
// }
