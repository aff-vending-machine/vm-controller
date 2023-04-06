package payment_channel

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/module/repository"
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
