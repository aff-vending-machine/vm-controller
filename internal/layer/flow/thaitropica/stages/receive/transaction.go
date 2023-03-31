package receive

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) transaction_refund(c *flow.Ctx) {
	ts := time.Now()
	totalReceived := c.Data.TotalReceived()
	totalPay := c.Data.TotalPay()
	// refund := c.Data.TotalPrice() - totalPay

	err := s.updateTransaction(c, map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE_BROKEN,
		"refund_at":         ts,
		"received_item_at":  ts,
		"received_quantity": totalReceived,
		"refund_price":      0,
		"paid_price":        totalPay,
	})
	if err != nil {
		log.Error().Err(err).Msg("unable to update transaction")
		return
	}
}

func (s *stageImpl) transaction_done(c *flow.Ctx) {
	ts := time.Now()
	totalReceived := c.Data.TotalReceived()
	totalPay := c.Data.TotalPay()
	err := s.updateTransaction(c, map[string]interface{}{
		"order_status":      enum.ORDER_STATUS_DONE,
		"received_item_at":  ts,
		"received_quantity": totalReceived,
		"paid_price":        totalPay,
		"is_error":          false,
		"error":             nil,
		"error_at":          nil,
	})
	if err != nil {
		log.Error().Err(err).Msg("unable to update transaction")
		return
	}
}

func (s *stageImpl) transaction_error(c *flow.Ctx, err_ error) {
	ts := time.Now()
	msg := err_.Error()

	err := s.updateTransaction(c, map[string]interface{}{
		"is_error": true,
		"error":    msg,
		"error_at": ts,
	})
	if err != nil {
		log.Error().Err(err).Msg("unable to update transaction")
		return
	}
}

func (s *stageImpl) updateTransaction(c *flow.Ctx, data map[string]interface{}) error {
	filter := makeMerchantOrderID(c.Data.MerchantOrderID)

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	return nil
}
