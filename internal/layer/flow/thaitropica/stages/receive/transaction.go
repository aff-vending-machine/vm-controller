package receive

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateBrokenTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE_BROKEN,
		"refund_at":         time.Now(),
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"refund_price":      0,
		"paid_price":        c.Data.TotalPay(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateDoneTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE,
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"paid_price":        c.Data.TotalPay(),
		"is_error":          false,
		"error":             nil,
		"error_at":          nil,
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateErrorTransaction(c *flow.Ctx, err error) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"is_error": true,
		"error":    err.Error(),
		"error_at": time.Now(),
	}

	_, err = s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateCancelTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_CANCELLED,
		"cancelled_by":      "user",
		"cancelled_at":      time.Now(),
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"paid_price":        c.Data.TotalPay(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}
