package slot

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/infra/storage/sqlite/service"
	"vm-controller/internal/core/interface/slot"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	*service.RepositoryImpl[entity.Slot]
}

func New(db *gorm.DB) slot.Repository {
	db.AutoMigrate(&entity.Slot{})
	return &repositoryImpl{
		service.New[entity.Slot](db),
	}
}
