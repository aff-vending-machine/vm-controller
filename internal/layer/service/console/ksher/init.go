package ksher_console

type consoleImpl struct {
	retry int
}

func New() *consoleImpl {
	return &consoleImpl{3}
}
