package payment

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/link2500"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) creditcard(c *flow.Ctx) {
	ctx, fn := context.WithCancel(c.UserCtx)
	s.CancelFn = fn

	req := link2500.SaleRequest{
		MerchantID: c.PaymentChannel.MerchantID,
		Price: c.Data.TotalPrice(),
	}
	res, err := s.link2500.Sale(ctx, c.PaymentChannel, &req)
	if c.Stage != "payment" || c.PaymentChannel.Channel != "creditcard" {
		log.Error().Str("stage", c.Stage).Str("channel", c.PaymentChannel.Channel).Msg("cancelled by user")
		return
	}

	if err != nil {
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
