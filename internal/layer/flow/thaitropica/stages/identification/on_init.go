package identification

import (
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/enum"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) OnInit(c *flow.Ctx) {
	ts := time.Now()
	c.Data.MerchantOrderID = ts.Format("20060102150405") + c.PaymentChannel.MerchantID
	err := s.transactionRepo.InsertOne(
		c.UserCtx,
		&entity.Transaction{
			MerchantOrderID:     c.Data.MerchantOrderID,
			MachineSerialNumber: c.Machine.SerialNumber,
			Location:            c.Machine.Location,
			RawCart:             c.Data.Raw(),
			OrderQuantity:       c.Data.TotalQuantity(),
			OrderPrice:          c.Data.TotalPrice(),
			OrderStatus:         enum.ORDER_STATUS_ORDERED,
			OrderedAt:           ts,
			PaymentChannel:      "free",
			PaymentRequestedAt:  &ts,
			RefundPrice:         0,
			ReceivedQuantity:    0,
			PaidPrice:           0,
			IsError:             false,
		})
	if err != nil {
		log.Error().Err(err).Msg("unable to create transaction")
		s.ui.SendError(c.UserCtx, "identification", err.Error())
		return
	}

	s.stacks = make(map[string]*OTPStack)
	s.ui.SendMailRequest(c.UserCtx, c.Data.MerchantOrderID)
	// waiting for email
	s.console_processing(c)
}
