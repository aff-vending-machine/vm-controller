package jetts

import (
	"context"
	"strings"
	"time"
	"vm-controller/internal/core/flow"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) ListenEvent(sn string) {
	ctx := context.Background()
	timeout := 2 * time.Minute
	uc.context.UserCtx = ctx

	uc.stages[flow.IDLE_STAGE].OnInit(uc.context)

	uc.watchdog = time.NewTicker(timeout)
	defer uc.watchdog.Stop()

	go func() {
		for {
			uc.lookup(ctx, timeout)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			uc.event(ctx, "EVENT")
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (uc *Flow) lookup(ctx context.Context, timeout time.Duration) {
	select {
	case <-uc.watchdog.C:
		if uc.context.Stage != flow.IDLE_STAGE {
			uc.context.Stage = flow.IDLE_STAGE
			uc.OnInit(ctx)
		}

	case <-uc.context.ClearWatchdog:
		uc.watchdog.Reset(timeout)

	case stage := <-uc.context.ChangeStage:
		log.Debug().Str("stage", string(stage)).Msg("stage changed")
		uc.context.Stage = stage
		uc.OnInit(ctx)

		if stage == flow.IDLE_STAGE || stage == flow.PAYMENT_STAGE || stage == flow.RECEIVE_STAGE {
			uc.watchdog.Stop()
		} else {
			uc.watchdog.Reset(timeout)
		}
	}
}

func (uc *Flow) event(ctx context.Context, key string) {
	event, err := uc.queueHw.Pop(ctx, key)
	if err != nil {
		if !strings.Contains(err.Error(), "redis: nil") {
			log.Error().Str("key", key).Err(err).Msg("Queue is not available")
			return
		}
	}
	if event == nil {
		return
	}

	switch event.Status {
	case "G0":
		log.Info().Str("key", event.ToValueCode()).Msg("Gate Closed")
		// uc.machineTp.IsClosed(ctx, uc.context.Machine)

	case "G1":
		log.Info().Str("key", event.ToValueCode()).Msg("Gate Opened")
		// uc.machineTp.IsOpened(ctx, uc.context.Machine)

	case "Z0": // reset
		uc.context.ChangeStage <- flow.IDLE_STAGE

	case "Z1": // emergency
		uc.context.ChangeStage <- flow.EMERGENCY_STAGE

	case "Z2": // open-gate
		uc.queueHw.PushCommand(ctx, "COMMAND", "OPEN_GATE")

	default:
		log.Error().Str("key", event.ToValueCode()).Msg("Unknown Key")
	}
}
