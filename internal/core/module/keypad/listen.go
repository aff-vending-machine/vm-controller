package keypad

import (
	"time"
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
