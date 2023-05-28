package machine

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/infra/storage/sqlite/service"
	"vm-controller/internal/core/interface/machine"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	*service.RepositoryImpl[entity.Machine]
}

func New(db *gorm.DB) machine.Repository {
	db.AutoMigrate(&entity.Machine{})
	return &repositoryImpl{
		service.New[entity.Machine](db),
	}
}
