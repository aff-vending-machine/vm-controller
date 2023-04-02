package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
	"github.com/rs/zerolog/log"
)

var mem = make(map[string][]hardware.Event, 0)

func (*consoleImpl) Push(ctx context.Context, key string, data hardware.Event) error {
	_, span := trace.Start(ctx)
	defer span.End()

	fmt.Printf("Push %s: %s \n", key, data.ToValueCode())
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
		fmt.Printf("no %s in memory \n", key)
	} else if len(event) > 0 {
		fmt.Printf("Pop %s: %s \n", key, event[0].ToValueCode())
		mem[key] = mem[key][1:]
	} else {
		fmt.Printf("Pop %s: <EMPTY> \n", key)
	}

	if len(event) == 0 {
		return nil, nil
	}

	return &event[0], nil
}

func (*consoleImpl) Clear(ctx context.Context) error {
	_, span := trace.Start(ctx)
	defer span.End()

	fmt.Printf("Clear Queue\n")
	mem = make(map[string][]hardware.Event, 0)
	return nil
}

func (*consoleImpl) PushCommand(ctx context.Context, key string, val string) error {
	_, span := trace.Start(ctx)
	defer span.End()

	fmt.Printf("Push %s: %s \n", key, val)
	if mem[key] == nil {
		mem[key] = make([]hardware.Event, 0)
	}
	mem[key] = append(mem[key], hardware.Event{Status: val})
	return nil
}

func (*consoleImpl) Polling(ctx context.Context, key string, total int, handler hardware.QueueHandler) {
	_, span := trace.Start(ctx)
	defer span.End()

	fmt.Printf("Polling: %s\n", key)
	for k, v := range mem["QUEUE"] {
		time.Sleep(5 * time.Second)
		v.Status = "S0"
		fmt.Printf("Process %d: %s \n", k, v.ToValueCode())
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

	fmt.Printf("Clear Stack\n")
	mem = make(map[string][]hardware.Event, 0)
}
