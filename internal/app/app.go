package app

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/app/preload"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/app/registry"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/app/router/fiber"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/app/router/ws"
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
		preload.InitAliPay(uc.PaymentChannel)
		preload.InitPromptPay(uc.PaymentChannel)
		preload.InitWechatPay(uc.PaymentChannel)
		preload.InitCreditCard(uc.PaymentChannel)
	}

	ws.New(cfg.WebSocket).Serve(dn.WebSocket, dr.WebSocket)
	fiber.New(cfg.Fiber).Serve(cfg.Fiber.Port, dr.HTTP)
	fw.ThaiTropica.ListenEvent(context.Background(), sn)

	log.Debug().Msg("start application")
}
