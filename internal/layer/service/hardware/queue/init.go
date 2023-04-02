package queue

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/go-redis/redis/v8"
)

type hardwareImpl struct {
	client *redis.Client
	stacks map[string]*hardware.Event
}

func New(conf config.RedisConfig) *hardwareImpl {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       0, // use default DB
	})
	boot.AddCloseFn(rdb.Close)

	err := rdb.Ping(context.Background()).Err()
	boot.TerminateWhenError(err)

	return &hardwareImpl{rdb, make(map[string]*hardware.Event)}
}
