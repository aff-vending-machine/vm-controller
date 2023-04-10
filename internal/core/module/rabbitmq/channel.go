package rabbitmq

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

// Connection amqp.Connection wrapper
type Connection struct {
	*amqp091.Connection
}

// Channel wrap amqp.Connection.Channel, get a auto reconnect channel
func (c *Connection) Channel() (*Channel, error) {
	mutex.Lock()
	defer mutex.Unlock()

	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}

	channel := &Channel{Channel: ch}

	go func() {
		for {
			reason, ok := <-channel.Channel.NotifyClose(make(chan *amqp091.Error))
			// exit this goroutine if closed by developer
			if !ok || channel.IsClosed() {
				// log.Debug().Msg("channel closed by developer")
				channel.Close() // close again, ensure closed flag set when connection closed
				break
			}
			log.Warn().Err(reason).Msg("channel closed")

			// reconnect if not closed by developer
			for {
				// wait for connection reconnect
				if c.Connection.IsClosed() {
					time.Sleep(time.Second)
					continue
				}

				ch, err := c.Connection.Channel()
				if err == nil {
					log.Debug().Msg("channel recreate success")
					channel.Channel = ch
					break
				}

				log.Error().Err(err).Msg("unable to recreate channel")
				time.Sleep(delay * time.Second)
			}
		}
	}()

	return channel, nil
}
