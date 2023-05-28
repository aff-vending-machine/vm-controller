package slot

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/interface/slot"
)

type usecaseImpl struct {
	slotRepo slot.Repository
}

func New(r slot.Repository) *usecaseImpl {
	return &usecaseImpl{r}
}
