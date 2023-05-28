package payment_channel

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/infra/storage/sqlite/service"
	"vm-controller/internal/core/interface/payment_channel"

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
