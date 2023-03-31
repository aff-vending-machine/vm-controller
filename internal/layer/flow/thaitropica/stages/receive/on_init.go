package receive

import (
	"fmt"
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	s.queue.Clear(c.UserCtx)

	if err := s.addEvents(c); err != nil {
		log.Error().Err(err).Msg("unable to add events")
		return
	}

	s.console(c)

	go s.checkEvent(c)
}

func (s *stageImpl) checkEvent(c *flow.Ctx) {
	total := c.Data.TotalQuantity()

	log.Info().Msg("start polling")
	s.polling = true
	s.queue.Polling(c.UserCtx, "RESPONSE", total, s.feedback(c))
	s.polling = false
	log.Info().Msg("stop polling")

	if c.Data.TotalQuantity() != c.Data.TotalReceived() {
		s.transaction_refund(c)
		s.void(c)
	} else {
		s.transaction_done(c)
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

func (s *stageImpl) void(c *flow.Ctx) {
	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "creditcard":
		if c.Data.TotalReceived() == 0 {
			s.smartedc.Void(c.UserCtx, &smartedc.VoidRequest{
				TradeType:        "CARD",
				InvoiceNo:        c.Data.InvoiceNo,
				CardApprovalCode: c.Data.CardApprovalCode,
				Amount:           c.Data.TotalPrice(),
				TransactionType:  "VOID",
				POSRefNo:         c.Data.MerchantOrderID,
			})
		}
		return
	default:
	}
}
