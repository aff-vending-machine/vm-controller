package receive

import (
	"time"

	"vm-controller/internal/core/domain/enum"
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateMachineFailedTransaction(c *flow.Ctx) error {
	ts := time.Now()
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"raw_card":          c.Data.Raw(),
		"order_status":      enum.ORDER_STATUS_DONE_BROKEN,
		"refund_at":         ts,
		"received_item_at":  ts,
		"received_quantity": c.Data.TotalReceived(),
		"refund_price":      0,
		"paid_price":        c.Data.TotalPay(),
	}
	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateDoneTransaction(c *flow.Ctx) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"raw_card":          c.Data.Raw(),
		"order_status":      enum.ORDER_STATUS_DONE,
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"paid_price":        c.Data.TotalPay(),
		"is_error":          false,
		"error":             nil,
		"error_at":          nil,
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateErrorTransaction(c *flow.Ctx, err error) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"raw_card": c.Data.Raw(),
		"is_error": true,
		"error":    err.Error(),
		"error_at": time.Now(),
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateCancelTransaction(c *flow.Ctx) {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"raw_card":          c.Data.Raw(),
		"order_status":      enum.ORDER_STATUS_CANCELLED,
		"cancelled_by":      "user",
		"cancelled_at":      time.Now(),
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"paid_price":        c.Data.TotalPay(),
	}

	_, err := s.transactionRepo.Update(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}
