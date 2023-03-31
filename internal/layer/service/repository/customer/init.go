package customer_repository

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Customer]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Customer](db)
	db.AutoMigrate(&entity.Customer{})
	return &repositoryImpl{based}
}
