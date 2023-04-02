package payment

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) creditcard(c *flow.Ctx) {
	ctx, fn := context.WithCancel(c.UserCtx)
	s.CancelFn = fn

	s.showCreditCard(c)

	req := link2500.SaleRequest{
		Price: c.Data.TotalPrice(),
	}
	res, err := s.link2500.Sale(ctx, c.PaymentChannel, &req)
	if c.Stage != "payment" || c.PaymentChannel.Name != "creditcard" {
		log.Error().Msg("cancelled by user")
		return
	}

	if err != nil {
		s.error(c, err, "creditcard is out of service")
		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- "emergency"
			return
		}

		c.ChangeStage <- "payment_channel"
		return
	}

	if res.ResponseText != "APPROVED" {
		err = fmt.Errorf("%s: %s", res.InvoiceNumber, res.ResponseText)
		s.error(c, err, "creditcard is rejected")
		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- "emergency"
			return
		}

		c.ChangeStage <- "payment_channel"
		return
	}

	err = s.updatePaidTransaction(c)
	if err != nil {
		c.ChangeStage <- "emergency"
		return
	}

	c.ChangeStage <- "receive"
}
