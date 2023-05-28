package jetts

import (
	"vm-controller/configs"
	"vm-controller/internal/boot/app/jetts/registry"
	"vm-controller/internal/boot/migrator"
	"vm-controller/internal/boot/router/rpc"
	"vm-controller/internal/boot/router/websocket"

	"github.com/rs/zerolog/log"
)

func Run(cfg configs.Config) {
	log.Debug().Msg("init application")

	var (
		infra     = registry.NewInfrastructure(cfg)
		service   = registry.NewService(infra)
		usecase   = registry.NewUsecase(service)
		flow      = registry.NewFlow(service)
		transport = registry.NewTransport(usecase, flow)
	)

	machine := migrator.InitMachine(cfg.App, usecase.Machine)

	if cfg.App.Preload {
		log.Debug().Msg("preload")
		migrator.InitMachineSlot(usecase.Slot)
		migrator.InitTestPay(usecase.PaymentChannel)
		migrator.InitPromptPay(usecase.PaymentChannel)
		migrator.InitCreditCard(usecase.PaymentChannel)
	}

	rpc.New(infra.RabbitMQ).Serve(machine.SerialNumber, transport.RPC)
	websocket.New(infra.WebSocket).Serve(service.WebSocket, transport.WebSocket)
	flow.ListenEvent(machine.SerialNumber)

	log.Debug().Msg("start application")
}
