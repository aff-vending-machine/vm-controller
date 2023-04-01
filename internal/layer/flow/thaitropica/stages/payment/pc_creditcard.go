package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/smartedc"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) creditcard(c *flow.Ctx) {
	ts := time.Now()
	c.Data.MerchantOrderID = ts.Format("20060102150405")
	ctx, fn := context.WithCancel(c.UserCtx)
	s.CancelFn = fn

	err := s.transactionRepo.InsertOne(
		c.UserCtx,
		&entity.Transaction{
			MerchantOrderID:     c.Data.MerchantOrderID,
			MachineSerialNumber: c.Machine.SerialNumber,
			Location:            c.Machine.Location,
			RawCart:             c.Data.Raw(),
			OrderQuantity:       c.Data.TotalQuantity(),
			OrderPrice:          c.Data.TotalPrice(),
			OrderStatus:         enum.ORDER_STATUS_ORDERED,
			OrderedAt:           ts,
			PaymentChannel:      c.PaymentChannel.Channel,
			PaymentRequestedAt:  &ts,
			RefundPrice:         0,
			ReceivedQuantity:    0,
			PaidPrice:           0,
			IsError:             false,
		})
	if err != nil {
		log.Error().Err(err).Msg("unable to create transaction")
		s.ui.SendError(c.UserCtx, "payment", err.Error())
		c.ChangeStage <- "payment_channel"
		return
	}

	if c.Stage != "payment" {
		log.Error().Msg("cancelled by user")
		return
	}

	req := smartedc.SaleRequest{
		TradeType:       "CARD",
		Amount:          c.Data.TotalPrice(),
		TransactionType: "SALE",
		POSRefNo:        c.Data.MerchantOrderID,
	}
	res, err := s.smartedc.Sale(ctx, &req)
	if c.Stage != "payment" {
		log.Error().Msg("cancelled by user")
		return
	}
	if err != nil {
		log.Error().Interface("request", req).Interface("response", res).Err(err).Msg("unable to process payment")
		s.ui.SendError(c.UserCtx, "payment", "creditcard is out of service")
		c.ChangeStage <- "payment_channel"
		return
	}

	if res.POSRefNo != c.Data.MerchantOrderID {
		log.Error().Interface("request", req).Interface("response", res).Str("POSRefNo", res.POSRefNo).Str("MerchantOrderID", c.Data.MerchantOrderID).Msg("POS reference number is not matched")
		s.ui.SendError(c.UserCtx, "payment", "POS reference number is not matched")
		c.ChangeStage <- "payment_channel"
		return
	}

	if res.ResponseMsg == "SUCCESS" {
		log.Info().Msg("paid")
		ts := time.Now()
		filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
		_, errx := s.transactionRepo.UpdateMany(
			c.UserCtx,
			filter,
			map[string]interface{}{
				"order_status":      enum.ORDER_STATUS_PAID,
				"confirmed_paid_by": "machine",
				"confirmed_paid_at": ts,
				"reference1":        res.InvoiceNo,
				"reference2":        res.CardApprovalCode,
				"reference3":        res.CardNo,
			})
		if errx != nil {
			log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Str("onPaid", "PAID").Msg("TRANSACTION: unable to update transaction")
		}

		c.Data.InvoiceNo = res.InvoiceNo
		c.Data.CardApprovalCode = res.CardApprovalCode

		s.ui.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
		c.ChangeStage <- "receive"
		return
	}

	log.Error().Interface("request", req).Interface("response", res).Msg("payment is not success")
	s.ui.SendError(c.UserCtx, "payment", res.ResponseMsg)
	c.ChangeStage <- "payment_channel"
}
