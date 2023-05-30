package payment

import (
	"context"
	"fmt"
	"time"

	"vm-controller/internal/core/domain/ksher"
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/conv"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) promptpay(c *flow.Ctx) {
	ctx, fn := context.WithCancel(c.UserCtx)
	s.CancelFn = fn

	req := ksher.CreateOrderBody{
		MerchantOrderID: c.Data.MerchantOrderID,
		Amount:          int(c.Data.TotalPrice() * 100),
		Timestamp:       fmt.Sprintf("%d", time.Now().Unix()),
		Channel:         ksher.PROMPTPAY,
		DeviceID:        c.Machine.Name,
	}
	res, err := s.ksher.CreateOrder(ctx, c.PaymentChannel, &req)
	if c.Stage != flow.PAYMENT_STAGE || c.PaymentChannel.Channel != "promptpay" {
		log.Error().Str("stage", string(c.Stage)).Str("channel", c.PaymentChannel.Channel).Msg("cancelled by user")
		return
	}

	if err != nil {
		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, err.Error())

		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		c.ChangeStage <- flow.CHANNEL_STAGE
		return
	}

	if res.ErrorCode != ksher.SUCCESS {
		err = fmt.Errorf("%s: %s", res.ErrorCode, res.ErrorMessage)
		s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, err.Error())

		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- flow.EMERGENCY_STAGE
			return
		}

		c.ChangeStage <- flow.CHANNEL_STAGE
		return
	}

	// short
	res.Reserved1 = res.Reserved1[:10] + "..."
	raw, err := conv.StructToString(res)
	if err != nil {
		log.Error().Err(err).Interface("response", res).Msg("unable to convert struct to string")
	}

	err = s.updateReferenceTransaction(c, res.Reference, res.GatewayOrderID, res.AcquirerOrderID, raw)
	if err != nil {
		c.ChangeStage <- flow.EMERGENCY_STAGE
		return
	}
	s.frontendWs.SendQRCode(c.UserCtx, c.Data.MerchantOrderID, res.Reference, c.Data.TotalQuantity(), c.Data.TotalPrice())

	go s.pollingPromptpay(c, ctx, req.Timestamp)
}

func (s *stageImpl) pollingPromptpay(c *flow.Ctx, ctx context.Context, timestamp string) {
	s.ticker = time.NewTicker(5 * time.Second)
	defer s.ticker.Stop()

	for {
		select {
		case <-s.ticker.C:
			if c.Stage != flow.PAYMENT_STAGE || c.PaymentChannel.Channel != "promptpay" {
				log.Error().Msg("cancelled by user")
				return
			}

			req := ksher.CheckOrderQuery{
				Timestamp: timestamp,
			}
			res, err := s.ksher.CheckOrder(ctx, c.PaymentChannel, c.Data.MerchantOrderID, &req)
			if err != nil {
				err = fmt.Errorf("%s: %s", res.ErrorCode, res.ErrorMessage)
				s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, err.Error())

				err = s.updateErrorTransaction(c, err)
				if err != nil {
					c.ChangeStage <- flow.EMERGENCY_STAGE
					return
				}

				c.ChangeStage <- flow.CHANNEL_STAGE
				return
			}

			if res.ErrorCode == ksher.SUCCESS {
				err = s.updatePaidTransaction(c)
				if err != nil {
					c.ChangeStage <- flow.EMERGENCY_STAGE
					return
				}

				s.frontendWs.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
				c.ChangeStage <- flow.RECEIVE_STAGE
				return
			}

		case <-time.After(5 * time.Minute):
			if c.Stage != flow.PAYMENT_STAGE || c.PaymentChannel.Channel != "promptpay" {
				log.Error().Msg("cancelled by user")
				return
			}

			s.frontendWs.SendError(c.UserCtx, flow.PAYMENT_STAGE, "timeout")

			err := s.updateCancelTransaction(c, "machine")
			if err != nil {
				c.ChangeStage <- flow.EMERGENCY_STAGE
				return
			}

			c.ChangeStage <- flow.CHANNEL_STAGE
			return
		}

		time.Sleep(100 * time.Millisecond)

		if s.bypass {
			err := s.updateReferenceTransaction(c, "BYPASS", "", "", "BYPASS")
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
	}

}
