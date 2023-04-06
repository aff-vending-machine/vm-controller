package slot

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/internal/layer/service/repository"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
)

type usecaseImpl struct {
	slotRepo repository.Slot
}

func New(r repository.Slot) *usecaseImpl {
	return &usecaseImpl{r}
}

func makeFilter(req *request.Filter) []string {
	filter := []string{}

	if req.Offset != nil {
		filter = append(filter, fmt.Sprintf(":OFFSET:%d", *req.Offset))
	}
	if req.Limit != nil {
		filter = append(filter, fmt.Sprintf(":LIMIT:%d", *req.Limit))
	}

	return filter
}
