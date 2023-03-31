package slot_usecase

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/service/repository"
)

type usecaseImpl struct {
	slotRepo repository.Slot
}

func New(r repository.Slot) *usecaseImpl {
	return &usecaseImpl{r}
}
