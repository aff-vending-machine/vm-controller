package payment

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updatePaidTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_PAID,
		"confirmed_paid_by": "machine",
		"confirmed_paid_at": time.Now(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updatePaidTransactionWithRef(c *flow.Ctx, ref1 string, ref2 string, ref3 string) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_PAID,
		"confirmed_paid_by": "machine",
		"confirmed_paid_at": time.Now(),
		"reference1":        ref1,
		"reference2":        ref2,
		"reference3":        ref3,
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateCancelTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"cancelled_by": "user",
		"cancelled_at": time.Now(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateCancelTransactionByMachine(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"cancelled_by": "machine",
		"cancelled_at": time.Now(),
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
