package app

import (
	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/internal/boot/preload"
	"github.com/aff-vending-machine/vm-controller/internal/boot/registry"
	"github.com/aff-vending-machine/vm-controller/internal/boot/router/rpc"
	"github.com/aff-vending-machine/vm-controller/internal/boot/router/websocket"
	"github.com/rs/zerolog/log"
)

func Run(cfg configs.Config) {
	var (
		infra     = NewInfrastructure(cfg)
		service   = NewService(infra)
		usecase   = NewUsecase(service)
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

	rpc.New(infra.RabbitMQ).Serve(machine.SerialNumber, transport.RPC)
	websocket.New(infra.WebSocket).Serve(service.WebSocket, transport.WebSocket)
	flow.Jetts.ListenEvent(machine.SerialNumber)

	log.Debug().Msg("start application")
}
