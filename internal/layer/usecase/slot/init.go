package slot

import (
	"vm-controller/internal/core/interface/slot"
)

type usecaseImpl struct {
	slotRepo slot.Repository
}

func New(r slot.Repository) *usecaseImpl {
	return &usecaseImpl{r}
}
