package ksher

type consoleImpl struct {
	retry int
}

func New() *consoleImpl {
	return &consoleImpl{3}
}
