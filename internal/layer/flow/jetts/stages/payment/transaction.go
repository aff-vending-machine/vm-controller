package payment

import (
	"time"

	"vm-controller/internal/core/domain/enum"
	"vm-controller/internal/core/flow"
	"vm-controller/pkg/helpers/db"

	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateReferenceTransaction(c *flow.Ctx, ref1 string, ref2 string, ref3 string, raw string) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"reference1":    ref1,
		"reference2":    ref2,
		"reference3":    ref3,
		"raw_reference": raw,
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updatePaidTransaction(c *flow.Ctx) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_PAID,
		"confirmed_paid_by": "machine",
		"confirmed_paid_at": time.Now(),
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

func (s *stageImpl) updateTestTransaction(c *flow.Ctx) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_PAID,
		"confirmed_paid_by": "test",
		"confirmed_paid_at": time.Now(),
		"is_error":          false,
		"error":             "(test)",
		"error_at":          time.Now(),
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}

func (s *stageImpl) updateCancelTransaction(c *flow.Ctx, by string) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"cancelled_by": by,
		"cancelled_at": time.Now(),
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
	return nil
}

func (s *stageImpl) updateErrorTransaction(c *flow.Ctx, err error) error {
	filter := db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"is_error":     true,
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"error":        err.Error(),
		"error_at":     time.Now(),
	}

	_, errx := s.transactionRepo.Update(c.UserCtx, filter, data)
	if errx != nil {
		log.Error().Err(errx).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
		return errx
	}

	return nil
}
