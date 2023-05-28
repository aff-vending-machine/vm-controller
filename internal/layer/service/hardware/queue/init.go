package queue

import (
	"vm-controller/internal/core/domain/hardware"

	"github.com/go-redis/redis/v8"
)

type hardwareImpl struct {
	client *redis.Client
	stacks map[string]*hardware.Event
}

func New(client *redis.Client) *hardwareImpl {

	return &hardwareImpl{
		client: client,
		stacks: make(map[string]*hardware.Event),
	}
}
