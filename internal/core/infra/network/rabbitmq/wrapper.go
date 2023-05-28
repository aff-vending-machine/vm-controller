package rabbitmq

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
)

type Wrapper struct {
	*Connection
}

func New(cfg configs.RabbitMQConfig) *Wrapper {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		cfg.Protocol,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Path,
	)
	conn, err := Dial(url)
	boot.TerminateWhenError(err)

	return &Wrapper{
		Connection: conn,
	}
}
