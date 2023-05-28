package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite/service"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/payment_channel"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*service.RepositoryImpl[entity.PaymentChannel]
}

func New(db *gorm.DB) payment_channel.Repository {
	db.AutoMigrate(&entity.PaymentChannel{})
	return &repositoryImpl{
		service.New[entity.PaymentChannel](db),
	}
}
