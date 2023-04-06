package rabbitmq

import (
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

const delay = 5 // reconnect after delay seconds
var mutex sync.Mutex
var stacks map[string]*Connection = make(map[string]*Connection)

// Dial wrap amqp.Dial, dial and get a reconnect connection
func Dial(url string) (*Connection, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if stacks[url] != nil {
		return stacks[url], nil
	}

	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	stacks[url] = &Connection{Connection: conn}

	go func() {
		for {
			reason, ok := <-stacks[url].Connection.NotifyClose(make(chan *amqp091.Error))
			// exit this goroutine if closed by developer
			if !ok {
				log.Debug().Msg("connection closed by developer")
				break
			}
			log.Warn().Err(reason).Msg("connection closed")

			// reconnect if not closed by developer
			for {
				// wait 1s for reconnect
				time.Sleep(delay * time.Second)

				conn, err := amqp091.Dial(url)
				if err == nil {
					stacks[url].Connection = conn
					log.Debug().Msg("reconnect success")
					break
				}

				log.Warn().Err(err).Msg("unable to reconnect")
			}
		}
	}()

	return stacks[url], nil
}
