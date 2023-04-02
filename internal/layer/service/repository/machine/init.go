package machine_repository

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/pkg/module/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Machine]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Machine](db)
	db.AutoMigrate(&entity.Machine{})
	return &repositoryImpl{based}
}
