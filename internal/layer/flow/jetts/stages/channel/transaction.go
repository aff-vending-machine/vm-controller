package channel

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/enum"
	"github.com/aff-vending-machine/vm-controller/internal/core/flow"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/errs"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) createTransaction(c *flow.Ctx) error {
	ts := time.Now()
	c.Data.MerchantOrderID = fmt.Sprintf("%s%s", c.Machine.Name, ts.Format("20060102150405"))

	_, err := s.transactionRepo.FindOne(c.UserCtx, db.NewQuery().AddWhere("merchant_order_id = ?", c.Data.MerchantOrderID))
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

	_, err = s.transactionRepo.Create(c.UserCtx, data)
	if err != nil {
		return err
	}

	return nil
}
