package queue

import (
	"context"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

var mem = make(map[string][]hardware.Event, 0)

func (*consoleImpl) Push(ctx context.Context, key string, data hardware.Event) error {
	_, span := trace.Start(ctx)
	defer span.End()

	log.Debug().Str("key", key).Str("event", data.ToValueCode()).Msg("EVENT: Push")
	if mem[key] == nil {
		mem[key] = make([]hardware.Event, 0)
	}
	mem[key] = append(mem[key], data)
	return nil
}

func (*consoleImpl) Pop(ctx context.Context, key string) (*hardware.Event, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	event := mem[key]
	if event == nil {
		//fmt.Printf("no %s in memory \n", key)
	} else if len(event) > 0 {
		log.Debug().Str("key", key).Str("event", event[0].ToValueCode()).Msg("EVENT: Pop")
		mem[key] = mem[key][1:]
	} else {
		log.Debug().Str("key", key).Msg("EVENT: Pop <empty>")
	}

	if len(event) == 0 {
		return nil, nil
	}

	return &event[0], nil
}

func (*consoleImpl) Clear(ctx context.Context) error {
	_, span := trace.Start(ctx)
	defer span.End()

	log.Debug().Msg("Clear QUEUE")
	mem = make(map[string][]hardware.Event, 0)
	return nil
}

func (*consoleImpl) PushCommand(ctx context.Context, key string, val string) error {
	_, span := trace.Start(ctx)
	defer span.End()

	log.Debug().Str("key", key).Str("command", val).Msg("EVENT: Push")
	if mem[key] == nil {
		mem[key] = make([]hardware.Event, 0)
	}
	mem[key] = append(mem[key], hardware.Event{Status: val})
	return nil
}

func (*consoleImpl) Polling(ctx context.Context, key string, total int, handler hardware.QueueHandler) {
	_, span := trace.Start(ctx)
	defer span.End()

	log.Debug().Str("key", key).Int("total", total).Msg("EVENT: Polling")
	for i, v := range mem["QUEUE"] { // take from QUEUE instead of RESPONSE
		time.Sleep(5 * time.Second)
		v.Status = "S0"
		log.Debug().Str("key", key).Int("index", i).Str("value", v.ToValueCode()).Msg("EVENT: Process")
		err := handler(&v)
		if err != nil {
			log.Error().Interface("event", v).Err(err).Msg("failed to handle queue polling")
		}
	}

	mem = make(map[string][]hardware.Event, 0)
}

func (*consoleImpl) Listen(ctx context.Context, key string, handler hardware.QueueHandler) {
}

func (*consoleImpl) ClearStack(ctx context.Context) {
	_, span := trace.Start(ctx)
	defer span.End()

	log.Debug().Msg("Clear STACK")
	mem = make(map[string][]hardware.Event, 0)
}
