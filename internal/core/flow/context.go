package flow

import (
	"context"

	"vm-controller/internal/core/domain/entity"
	"vm-controller/internal/core/domain/hardware"
)

type Ctx struct {
	Stage          Stage
	Data           *hardware.Data
	Machine        *entity.Machine
	PaymentChannel *entity.PaymentChannel
	Error          error
	Events         map[string]*hardware.Event
	ChangeStage    chan Stage
	ClearWatchdog  chan bool
	UserCtx        context.Context
}

type QueueHandler func(*hardware.Event) error

func NewContext() *Ctx {
	return &Ctx{
		Stage:          IDLE_STAGE,
		Data:           &hardware.Data{},
		Machine:        &entity.Machine{},
		PaymentChannel: &entity.PaymentChannel{},
		Error:          nil,
		Events:         make(map[string]*hardware.Event),
		ChangeStage:    make(chan Stage, 1),
		ClearWatchdog:  make(chan bool, 1),
	}
}
