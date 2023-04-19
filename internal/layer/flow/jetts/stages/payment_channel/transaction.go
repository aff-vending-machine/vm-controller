package payment_channel

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/enum"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/aff-vending-machine/vm-controller/pkg/errs"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) createTransaction(c *flow.Ctx) error {
	ts := time.Now()
	c.Data.MerchantOrderID = fmt.Sprintf("%s%s", c.Machine.Name, ts.Format("20060102150405"))

	_, err := s.transactionRepo.FindOne(c.UserCtx, []string{"merchant_order_id:=:%s", c.Data.MerchantOrderID})
	if errs.Not(err, "not found") {
		log.Error().Err(err).Msg("unable to create transaction")
		return err
	}
	if err == nil {
		return nil
	}

	data := &entity.Transaction{
		MerchantOrderID:     c.Data.MerchantOrderID,
		MachineSerialNumber: c.Machine.SerialNumber,
		Location:            c.Machine.Location,
		RawCart:             c.Data.Raw(),
		OrderQuantity:       c.Data.TotalQuantity(),
		OrderPrice:          c.Data.TotalPrice(),
		OrderedAt:           ts,
		OrderStatus:         enum.ORDER_STATUS_ORDERED,
		PaymentChannel:      c.PaymentChannel.Channel,
		PaymentRequestedAt:  &ts,
		RefundPrice:         0,
		ReceivedQuantity:    0,
		PaidPrice:           0,
		IsError:             false,
	}

	return s.transactionRepo.InsertOne(c.UserCtx, data)
}

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
