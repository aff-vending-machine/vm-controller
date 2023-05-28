package slot

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/repository"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	*repository.Template[entity.Slot]
}

func New(db *gorm.DB) *repositoryImpl {
	based := repository.New[entity.Slot](db)
	db.AutoMigrate(&entity.Slot{})
	return &repositoryImpl{based}
}
