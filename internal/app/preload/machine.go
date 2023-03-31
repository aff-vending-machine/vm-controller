package preload

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/layer/usecase/machine/request"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/boot"
)

func InitMachine(app config.AppConfig, uc machine.Usecase) string {
	ctx := context.TODO()

	sn, err := uc.StartUp(ctx,
		&request.StartUp{
			Codename:     "TTP",
			Name:         app.Machine.Name,
			SerialNumber: app.Machine.SerialNumber,
			Location:     app.Machine.Location,
			Type:         app.Machine.Type,
			Center:       app.Machine.Center,
			Vendor:       app.Machine.Vendor,
		})
	boot.TerminateWhenError(err)

	return sn
}
