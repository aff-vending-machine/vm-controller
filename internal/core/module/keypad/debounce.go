package keypad

import (
	"sync"
	"time"
)

type debouncerFunc func(func())

type debouncer struct {
	mu    sync.Mutex
	after time.Duration
	timer *time.Timer
}

func NewDebouncer(after time.Duration) debouncerFunc {
	d := &debouncer{after: after}

	return func(f func()) {
		d.add(f)
	}
}

func (d *debouncer) add(fn func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.after, fn)
}
