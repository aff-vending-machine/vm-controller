package payment_channel_repository

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.PaymentChannel]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.PaymentChannel](db)
	db.AutoMigrate(&entity.PaymentChannel{})
	return &repositoryImpl{based}
}
