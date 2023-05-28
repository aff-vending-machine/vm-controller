package slot

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/infra/storage/sqlite/service"
)

type Repository interface {
	service.Repository[entity.Slot]
}
