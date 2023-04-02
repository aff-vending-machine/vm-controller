package identification

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) updateReference(c *flow.Ctx, mail string, reference string, otp string) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"reference1": mail,
		"reference2": reference,
		"reference3": otp,
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		log.Error().Err(err).Str("order_id", c.Data.MerchantOrderID).Interface("data", data).Msg("TRANSACTION: unable to update transaction")
	}
}

func (s *stageImpl) updateIdentifiedTransaction(c *flow.Ctx) {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"confirmed_paid_by": "machine",
		"confirmed_paid_at": time.Now(),
		"order_status":      enum.ORDER_STATUS_IDENTIFIED,
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

func (s *stageImpl) updateEmailUsedTransaction(c *flow.Ctx, email string) error {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"error":    fmt.Sprintf("(email %s is used)", email),
		"error_at": time.Now(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	return nil
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
