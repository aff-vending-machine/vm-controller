package summary

import (
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/enum"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) createTransaction(c *flow.Ctx) error {
	data := &entity.Transaction{
		MerchantOrderID:     c.Data.MerchantOrderID,
		MachineSerialNumber: c.Machine.SerialNumber,
		Location:            c.Machine.Location,
		RawCart:             c.Data.Raw(),
		OrderQuantity:       c.Data.TotalQuantity(),
		OrderPrice:          c.Data.TotalPrice(),
		OrderedAt:           time.Now(),
		OrderStatus:         enum.ORDER_STATUS_ORDERED,
		RefundPrice:         0,
		ReceivedQuantity:    0,
		PaidPrice:           0,
		IsError:             false,
	}

	return s.transactionRepo.InsertOne(c.UserCtx, data)
}
