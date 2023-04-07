package keypad

import (
	"time"

	"github.com/rs/zerolog/log"
)

func (app *App) Listen() {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	var pkey string

	// Limit the keypress rate to 1 every 200 milliseconds.
	delay := 200 * time.Millisecond
	debounced := NewDebouncer(delay)

	for {
		row := polling(app.horizontal, app.vertical)
		col := polling(app.vertical, app.horizontal)

		if col < 0 || row < 0 {
			time.Sleep(delay)
			continue
		}

		debounced(func() {
			key := app.keys[row][col]
			if pkey == key {
				return
			}

			pkey = key
			log.Debug().Str("key", key).Bool("has handler", app.handler != nil).Msg("pressed")
			if app.handler != nil {
				(*app.handler)(key)
			}

			go func() {
				time.Sleep(delay)
				pkey = ""
			}()
		})

		time.Sleep(delay)
	}
}
