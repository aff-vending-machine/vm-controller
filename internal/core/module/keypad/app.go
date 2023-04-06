package keypad

import (
	"sync"

	"github.com/stianeikeland/go-rpio/v4"
)

type App struct {
	mutex      sync.Mutex
	keys       [][]string
	vertical   []rpio.Pin
	horizontal []rpio.Pin
	handler    *Handler
}

type Config struct {
	VerticalLine   []int
	HorizontalLine []int
}

type Handler func(key string) error
