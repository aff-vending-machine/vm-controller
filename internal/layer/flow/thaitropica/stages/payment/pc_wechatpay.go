package payment

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/lugentpay"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/utils"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) wechatpay(c *flow.Ctx) {
	ts := time.Now()
	accessCode := utils.GenerateRandom(8)
	req := lugentpay.QRCodeGenerateRequest{
		BillerCode:   c.PaymentChannel.BillerCode,
		BillerID:     c.PaymentChannel.BillerID,
		Reference1:   ts.Format("20060102150405"),
		Reference2:   c.PaymentChannel.MerchantID,
		Amount:       fmt.Sprintf("%0.02f", c.Data.TotalPrice()),
		StoreID:      c.PaymentChannel.StoreID,
		TerminalID:   c.PaymentChannel.TerminalID,
		MerchantName: c.PaymentChannel.MerchantName,
		AccessCode:   accessCode,
	}
	res, err := s.lugentpay.WechatPay(c.UserCtx, c.PaymentChannel, &req)
	if err != nil {
		log.Error().Interface("request", req).Interface("response", res).Err(err).Msg("unable to process payment")
		s.ui.SendError(c.UserCtx, "payment", err.Error())
		c.ChangeStage <- "payment_channel"
		return
	}

	if res.ResponseCode != "00" {
		log.Error().Interface("request", req).Interface("response", res).Msg("payment is not success")
		s.ui.SendError(c.UserCtx, "payment", res.ResponseDescription)
		c.ChangeStage <- "payment_channel"
		return
	}

	c.Data.MerchantOrderID = res.TransactionID
	s.qrcode = &res.QRCode

	s.ui.SendQRCode(c.UserCtx, c.Data.MerchantOrderID, res.QRCode, c.Data.TotalQuantity(), c.Data.TotalPrice())
	err = s.transactionRepo.InsertOne(
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
			Reference1:          &res.QRCode,
			Reference2:          &accessCode,
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

	go func(c *flow.Ctx) {
		time.Sleep(10 * time.Second)
		s.ticker = time.NewTicker(10 * time.Second)
		defer s.ticker.Stop()

		for {
			<-s.ticker.C
			if c.Stage != "payment" {
				return
			}

			req := &lugentpay.InquiryBody{
				TransactionID: c.Data.MerchantOrderID,
				AccessCode:    accessCode,
			}
			result, err := s.lugentpay.Inquiry(c.UserCtx, c.PaymentChannel, req)
			if err != nil {
				log.Error().Interface("request", req).Interface("response", result).Err(err).Msg("unable to check payment")
				s.updateErrorTransaction(c, err)
				s.ui.SendError(c.UserCtx, "payment", "unable to check order")
				c.ChangeStage <- "payment_channel"
				return
			}

			if result.ResponseCode == "00" {
				s.updatePaidTransaction(c)

				s.ui.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
				c.ChangeStage <- "receive"
				return
			}
		}
	}(c)
}
