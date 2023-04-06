package console

import (
	"bufio"
	"os"
	"time"
)

func (app *App) Listen() {
	for {
		reader := bufio.NewReader(os.Stdin)
		key, _, _ := reader.ReadRune()

		if app.handler != nil {
			(*app.handler)(string(key))
		}

		time.Sleep(100 * time.Millisecond)
	}
}
