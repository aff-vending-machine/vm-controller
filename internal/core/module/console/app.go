package console

type App struct {
	handler *Handler
}

type Handler func(key string) error
