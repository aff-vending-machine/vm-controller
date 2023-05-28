package redis

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/go-redis/redis/v8"
)

type Wrapper struct {
	*redis.Client
}

func New(conf configs.RedisConfig) *Wrapper {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       0, // use default DB
	})
	boot.AddCloseFn(client.Close)

	err := client.Ping(context.Background()).Err()
	boot.TerminateWhenError(err)

	return &Wrapper{
		client,
	}
}
