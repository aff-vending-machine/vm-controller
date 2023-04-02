package jetts

import (
	"context"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func (uc *Flow) ListenEvent(c context.Context, sn string) {
	uc.context.UserCtx = c

	timeout := 2 * time.Minute
	uc.stages["idle"].OnInit(uc.context)

	go func() {
		ctx := context.Background()
		uc.watchdog = time.NewTicker(timeout)
		defer uc.watchdog.Stop()

		for {
			uc.lookup(ctx, timeout)
			uc.event(ctx, "EVENT")
			time.Sleep(150 * time.Millisecond)
		}
	}()
}

func (uc *Flow) lookup(ctx context.Context, timeout time.Duration) {
	select {
	case <-uc.watchdog.C:
		log.Debug().Msg("watchdog timeout")
		uc.context.Stage = "idle"
		uc.OnInit(ctx)

	case <-uc.context.ClearWatchdog:
		uc.watchdog.Reset(timeout)

	case stage := <-uc.context.ChangeStage:
		log.Debug().Str("stage", stage).Msg("stage changed")
		uc.context.Stage = stage
		uc.OnInit(ctx)

		if stage == "done" {
			stage = "receive"
			uc.watchdog.Reset(timeout)
		} else if stage == "idle" || stage == "payment" || stage == "receive" {
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
		uc.context.ChangeStage <- "idle"

	case "Z1": // emergency
		uc.context.ChangeStage <- "emergency"

	case "Z2": // open-gate
		uc.queueHw.PushCommand(ctx, "COMMAND", "OPEN_GATE")

	default:
		log.Error().Str("key", event.ToValueCode()).Msg("Unknown Key")
	}
}
