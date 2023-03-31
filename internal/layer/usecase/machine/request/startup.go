package request

import (
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/utils"
)

type StartUp struct {
	Codename     string `json:"codename" validate:"required"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Location     string `json:"location,omitempty"`
	Type         string `json:"type,omitempty"`
	Center       string `json:"center"`
	Vendor       string `json:"vendor,omitempty"`
}

func (r *StartUp) ToEntity() *entity.Machine {
	if r.Name == "" {
		r.Name = strings.ToUpper(utils.GenerateRandom(6))
	}

	if r.SerialNumber == "" {
		r.SerialNumber = "VM-" + strings.ToUpper(r.Codename[:3]) + "-" + r.Name
	}

	if r.Center == "" {
		r.Center = "CT-APP-CENTER"
	}

	return &entity.Machine{
		Name:         r.Name,
		SerialNumber: r.SerialNumber,
		Location:     r.Location,
		Type:         r.Type,
		Vendor:       r.Vendor,
		Center:       r.Center,
		Status:       "out-of-service",
	}
}
