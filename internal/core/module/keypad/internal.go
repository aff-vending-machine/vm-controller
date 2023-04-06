package keypad

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func polling(cols, rows []rpio.Pin) int {
	input := -1

	for i := range cols {
		cols[i].Output()
		cols[i].Low()
		rows[i].Input()
		rows[i].PullDown()
	}

	for i := range cols {
		cols[i].High()
		for j := range rows {
			time.Sleep(100 * time.Nanosecond)
			logic := rows[j].Read()
			if logic == rpio.High {
				input = j
			}
		}
		time.Sleep(100 * time.Nanosecond)
		cols[i].Low()
	}

	return input
}
