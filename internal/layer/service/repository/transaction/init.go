package transaction

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite/service"
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/transaction"
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
