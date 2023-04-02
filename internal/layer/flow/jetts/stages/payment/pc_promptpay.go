package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/ksher"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
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
		DeviceID:        c.Machine.SerialNumber,
	}
	res, err := s.ksher.CreateOrder(ctx, c.PaymentChannel, &req)
	if c.Stage != "payment" || c.PaymentChannel.Name != "promptpay" {
		log.Error().Msg("cancelled by user")
		return
	}

	if err != nil {
		s.error(c, err, "promptpay is out of service")
		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- "emergency"
			return
		}

		c.ChangeStage <- "payment_channel"
		return
	}

	if res.ErrorCode != ksher.SUCCESS {
		err = fmt.Errorf("%s: %s", res.ErrorCode, res.ErrorMessage)
		s.error(c, err, "promptpay is rejected")
		err = s.updateErrorTransaction(c, err)
		if err != nil {
			c.ChangeStage <- "emergency"
			return
		}

		c.ChangeStage <- "payment_channel"
		return
	}

	err = s.updateReferenceTransaction(c, res.ID, res.AcquirerOrderID, res.GatewayOrderID)
	if err != nil {
		c.ChangeStage <- "emergency"
		return
	}
	s.showPromptPay(c, res)

	go s.pollingPromptpay(c, ctx, req.Timestamp)
}

func (s *stageImpl) pollingPromptpay(c *flow.Ctx, ctx context.Context, timestamp string) {
	s.ticker = time.NewTicker(5 * time.Second)
	defer s.ticker.Stop()

	for {
		select {
		case <-s.ticker.C:
			if c.Stage != "payment" || c.PaymentChannel.Name != "promptpay" {
				log.Error().Msg("cancelled by user")
				return
			}

			req := ksher.CheckOrderQuery{
				Timestamp: timestamp,
			}
			res, err := s.ksher.CheckOrder(ctx, c.PaymentChannel, c.Data.MerchantOrderID, &req)
			if err != nil {
				err = fmt.Errorf("%s: %s", res.ErrorCode, res.ErrorMessage)
				s.error(c, err, "promptpay is rejected")
				err = s.updateErrorTransaction(c, err)
				if err != nil {
					c.ChangeStage <- "emergency"
					return
				}

				c.ChangeStage <- "payment_channel"
				return
			}

			if res.ErrorCode == ksher.SUCCESS {
				err = s.updatePaidTransaction(c)
				if err != nil {
					c.ChangeStage <- "emergency"
					return
				}

				c.ChangeStage <- "receive"
				return
			}

		case <-time.After(5 * time.Minute):
			if c.Stage != "payment" || c.PaymentChannel.Name != "promptpay" {
				log.Error().Msg("cancelled by user")
				return
			}

			err := s.updateCancelTransaction(c)
			if err != nil {
				c.ChangeStage <- "emergency"
				return
			}

			c.ChangeStage <- "payment_channel"
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

}
