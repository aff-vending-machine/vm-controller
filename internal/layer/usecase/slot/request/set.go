package request

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/layer/usecase/slot/model"
)

type Set struct {
	Data []model.Slot `json:"data"`
}

func (r *Set) ToEntities() []entity.Slot {
	entities := make([]entity.Slot, len(r.Data))
	for i, slot := range r.Data {
		entities[i] = *slot.ToEntity()
	}

	return entities
}
