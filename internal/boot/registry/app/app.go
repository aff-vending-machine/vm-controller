package app

import (
	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/boot/preload"
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/boot/router/rpc"
	"github.com/aff-vending-machine/vm-controller/internal/boot/router/websocket"
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
		preload.InitTestPay(usecase.PaymentChannel)
		preload.InitPromptPay(usecase.PaymentChannel)
		preload.InitCreditCard(usecase.PaymentChannel)
	}

	rpc.New(module.RabbitMQ).Serve(machine.SerialNumber, transport.RPC)
	websocket.New(module.WebSocket).Serve(service.WebSocket, transport.WebSocket)
	flow.Jetts.ListenEvent(machine.SerialNumber)

	log.Debug().Msg("start application")
}
