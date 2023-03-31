package payment

import (
	"fmt"
	"strings"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	if !c.PaymentChannel.Active {
		s.ui.SendError(c.UserCtx, "payment", fmt.Sprintf("%s is out of service", c.PaymentChannel.Channel))
	}

	switch strings.ToLower(c.PaymentChannel.Channel) {
	case "promptpay":
		s.promptpay(c)
	case "creditcard":
		go s.creditcard(c)
	case "wechatpay":
		s.wechatpay(c)
	case "alipay":
		s.alipay(c)
	default:
		return
	}
}

func (s *stageImpl) onError(c *flow.Ctx, err error, msg string) {
	log.Error().Err(err).Msg(msg)
	ts := time.Now()
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	_, errx := s.transactionRepo.UpdateMany(
		c.UserCtx,
		filter,
		map[string]interface{}{
			"is_error": true,
			"error":    err.Error(),
			"error_at": ts,
		})
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Str("onError", msg).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) onPaid(c *flow.Ctx) {
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
		})
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Str("onPaid", "PAID").Msg("TRANSACTION: unable to update transaction")
	}

	s.ui.SendPaid(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
	c.ChangeStage <- "receive"
}

func (s *stageImpl) onCancel(c *flow.Ctx) {
	ts := time.Now()
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	_, errx := s.transactionRepo.UpdateMany(
		c.UserCtx,
		filter,
		map[string]interface{}{
			"order_status": enum.ORDER_STATUS_CANCELLED,
			"cancelled_by": "user",
			"cancelled_at": ts,
		})

	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Str("onCancel", "CANCELLED").Msg("TRANSACTION: unable to update transaction")
	}
}
