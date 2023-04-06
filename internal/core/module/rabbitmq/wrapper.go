package rabbitmq

import (
	"fmt"

	"github.com/aff-vending-machine/vm-controller/config"
	"github.com/rs/zerolog/log"
)

type Wrapper struct {
	*Connection
}

func New(cfg config.RabbitMQConfig) *Wrapper {
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
	if err != nil {
		log.Error().Str("url", url).Err(err).Msg("unable to dial RabbitMQ")
	}

	return &Wrapper{
		Connection: conn,
	}
}
