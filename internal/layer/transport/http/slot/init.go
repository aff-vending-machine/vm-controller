package slot_http

import "github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot"

type httpImpl struct {
	usecase slot.Usecase
}

func New(uc slot.Usecase) *httpImpl {
	return &httpImpl{uc}
}
