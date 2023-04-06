package response

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/machine/model"
)

type Machine = model.Machine

func ToMachine(e *entity.Machine) *Machine {
	return &Machine{
		Name:         e.Name,
		SerialNumber: e.SerialNumber,
		Branch:       e.Branch,
		Location:     e.Location,
		Type:         e.Type,
		Vendor:       e.Vendor,
		Center:       e.Center,
		Status:       e.Status,
	}
}
