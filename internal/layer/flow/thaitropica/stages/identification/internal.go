package identification

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
)

func makeMerchantOrderIDFilter(id string) []string {
	return []string{
		fmt.Sprintf("merchant_order_id:=:%s", id),
	}
}

func (s *stageImpl) updateReference(c *flow.Ctx, mail string, reference string, otp string) error {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"reference1": mail,
		"reference2": reference,
		"reference3": otp,
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *stageImpl) updateIdentified(c *flow.Ctx) error {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"confirmed_paid_by": "machine",
		"confirmed_paid_at": time.Now(),
		"order_status":      enum.ORDER_STATUS_IDENTIFIED,
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	s.ui.SendIdentified(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
	c.ChangeStage <- "receive"
	return nil
}

func (s *stageImpl) updateCancel(c *flow.Ctx) error {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"order_status": enum.ORDER_STATUS_CANCELLED,
		"cancelled_by": "machine",
		"cancelled_at": time.Now(),
	}

	_, err := s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	c.ChangeStage <- "order"
	return nil
}

func (s *stageImpl) updateError(c *flow.Ctx, err error) error {
	filter := makeMerchantOrderIDFilter(c.Data.MerchantOrderID)
	data := map[string]interface{}{
		"is_error": true,
		"error":    err.Error(),
		"error_at": time.Now(),
	}

	_, err = s.transactionRepo.UpdateMany(c.UserCtx, filter, data)
	if err != nil {
		return err
	}

	c.ChangeStage <- "order"
	return nil
}
