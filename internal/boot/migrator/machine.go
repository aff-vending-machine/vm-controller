package migrator

import (
	"context"

	"vm-controller/configs"
	"vm-controller/internal/core/interface/machine"
	"vm-controller/internal/layer/usecase/machine/request"
	"vm-controller/internal/layer/usecase/machine/response"
	"vm-controller/pkg/boot"
)

func InitMachine(app configs.AppConfig, uc machine.Usecase) *response.Machine {
	ctx := context.TODO()

	res, err := uc.StartUp(ctx,
		&request.StartUp{
			Codename:     "JET",
			Name:         app.Machine.Name,
			SerialNumber: app.Machine.SerialNumber,
			Location:     app.Machine.Location,
			Type:         app.Machine.Type,
			Center:       app.Machine.Center,
			Vendor:       app.Machine.Vendor,
		})
	boot.TerminateWhenError(err)

	return res
}
