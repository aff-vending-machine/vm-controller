package payment

import (
	"context"
	"fmt"
	"strings"

	"vm-controller/internal/core/domain/link2500"
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/conv"
	"vm-controller/pkg/helpers/errs"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) creditcard(c *flow.Ctx) {
	ctx, fn := context.WithCancel(c.UserCtx)
	s.CancelFn = fn

	req := link2500.SaleRequest{
		MerchantID: c.PaymentChannel.MerchantID,
		Price:      c.Data.TotalPrice(),
	}
	res, err := s.link2500.Sale(ctx, c.PaymentChannel, &req)
	if c.Stage != flow.PAYMENT_STAGE || c.PaymentChannel.Channel != "creditcard" {
		log.Error().Str("stage", string(c.Stage)).Str("channel", c.PaymentChannel.Channel).Msg("cancelled by user")
		return
	}

	if s.bypass {
		err = s.updateReferenceTransaction(c, "BYPASS", "", "", "BYPASS")
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		err = s.updatePaidTransaction(c)
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		s.frontendWs.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
		c.ChangeStage <- flow.RECEIVE_STAGE
		return
	}

	if err != nil {
		if errs.Is(err, "cancel") {
			if c.Stage == flow.PAYMENT_STAGE {
				c.Reset()
				c.ChangeStage <- flow.ORDER_STAGE
			}
			return
		}

		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, err.Error())

		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		c.ChangeStage <- flow.CHANNEL_STAGE
		return
	}

	if !strings.HasPrefix(res.ResponseText, "APPROVED") {
		err = fmt.Errorf("%s: %s", res.InvoiceNumber, res.ResponseText)
		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, err.Error())

		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		c.ChangeStage <- flow.CHANNEL_STAGE
		return
	}

	raw, err := conv.StructToString(res)
	if err != nil {
		log.Error().Err(err).Interface("response", res).Msg("unable to convert struct to string")
	}

	err = s.updateReferenceTransaction(c, res.BatchNumber+" - "+res.InvoiceNumber, res.CardIssuerName, res.PrimaryAccountNumber, raw)
	if err != nil {
		c.ChangeStage <- flow.EMERGENCY_STAGE
		return
	}

	err = s.updatePaidTransaction(c)
	if err != nil {
		c.ChangeStage <- flow.EMERGENCY_STAGE
		return
	}

	s.frontendWs.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
	c.ChangeStage <- flow.RECEIVE_STAGE
}
