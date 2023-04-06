package receive

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/enum"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateBrokenTransaction(c *flow.Ctx) error {
	ts := time.Now()
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE_BROKEN,
		"refund_at":         ts,
		"received_item_at":  ts,
		"received_quantity": c.Data.TotalReceived(),
		"refund_price":      0,
		"paid_price":        c.Data.TotalPay(),
	}
	_, errx := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateDoneTransaction(c *flow.Ctx) error {
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE,
		"received_item_at":  time.Now(),
		"received_quantity": c.Data.TotalReceived(),
		"paid_price":        c.Data.TotalPay(),
		"is_error":          false,
		"error":             nil,
		"error_at":          nil,
	}

	_, errx := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateErrorTransaction(c *flow.Ctx, err error) error {
	filter := []string{fmt.Sprintf("merchant_order_id:=:%s", c.Data.MerchantOrderID)}
	data := map[string]interface{}{
		"is_error": true,
		"error":    err.Error(),
		"error_at": time.Now(),
	}

	_, errx := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}
