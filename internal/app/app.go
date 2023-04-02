package app

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/app/preload"
	"github.com/aff-vending-machine/vm-controller/internal/app/registry"
	"github.com/aff-vending-machine/vm-controller/internal/app/router/fiber"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	var (
		dn = registry.NewAppDriven(cfg)
		uc = registry.NewAppUsecase(dn)
		fw = registry.NewAppFlow(dn)
		dr = registry.NewAppDriver(uc, fw)
	)

	sn := preload.InitMachine(cfg.App, uc.Machine)

	if cfg.App.Preload {
		log.Debug().Msg("preload")
		preload.InitMachineSlot(uc.Slot)
		preload.InitPromptPay(uc.PaymentChannel)
		preload.InitCreditCard(uc.PaymentChannel)
	}

	fiber.New(cfg.Fiber).Serve(cfg.Fiber.Port, dr.HTTP)
	fw.ThaiTropica.ListenEvent(context.Background(), sn)

	log.Debug().Msg("start application")
}
