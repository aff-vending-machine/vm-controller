package payment_channel

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/enum"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateTransaction(c *flow.Ctx, channel entity.PaymentChannel) error {
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	data := map[string]interface{}{
		"payment_channel":      channel.Channel,
		"payment_requested_at": time.Now(),
	}

	_, errx := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateCancelTransaction(c *flow.Ctx) error {
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	data := map[string]interface{}{
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"cancelled_by": "user",
		"cancelled_at": time.Now(),
	}

	_, errx := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
	return nil
}
