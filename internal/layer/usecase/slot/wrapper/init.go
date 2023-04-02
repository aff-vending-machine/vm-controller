package slot_wrapper

import (
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"
)

type wrapperImpl struct {
	usecase slot.Usecase
}

func New(uc slot.Usecase) *wrapperImpl {
	return &wrapperImpl{uc}
}
