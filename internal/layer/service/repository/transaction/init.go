package transaction

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/infra/storage/sqlite/service"
	"vm-controller/internal/core/interface/transaction"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	*service.RepositoryImpl[entity.Transaction]
}

func New(db *gorm.DB) transaction.Repository {
	db.AutoMigrate(&entity.Transaction{})
	return &repositoryImpl{
		service.New[entity.Transaction](db),
	}
}
