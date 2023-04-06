package app

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/boot/preload"
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/boot/router/keypad"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.BootConfig) {
	var (
		module    = NewModule(cfg)
		service   = NewService(module)
		usecase   = registry.NewUsecase(service)
		flow      = registry.NewFlow(service)
		transport = NewTransport(usecase, flow)
	)

	machine := preload.InitMachine(cfg.App, usecase.Machine)

	if cfg.App.Preload {
		log.Debug().Msg("preload")
		preload.InitMachineSlot(usecase.Slot)
		preload.InitPromptPay(usecase.PaymentChannel)
		preload.InitCreditCard(usecase.PaymentChannel)
	}

	keypad.New(module.Keypad).Scan(transport.Keypad)
	flow.Jetts.ListenEvent(machine.SerialNumber)

	log.Debug().Msg("start application")
}
