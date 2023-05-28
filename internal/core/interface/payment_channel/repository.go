package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite/service"
)

type Repository interface {
	service.Repository[entity.PaymentChannel]
}
