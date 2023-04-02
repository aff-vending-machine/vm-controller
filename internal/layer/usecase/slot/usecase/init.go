package slot_usecase

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
)

type usecaseImpl struct {
	slotRepo repository.Slot
}

func New(r repository.Slot) *usecaseImpl {
	return &usecaseImpl{r}
}
