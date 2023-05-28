package response

import (
	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/layer/usecase/slot/model"
)

type Slot = model.Slot

func ToSlotList(entities []entity.Slot) []model.Slot {
	results := make([]model.Slot, len(entities))
	for i, e := range entities {
		results[i] = *model.ToSlot(&e)
	}

	return results
}
