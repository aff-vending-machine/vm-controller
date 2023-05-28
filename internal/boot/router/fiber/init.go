package fiber

import "vm-controller/internal/core/infra/network/fiber"

type server struct {
	*fiber.Wrapper
}

func New(app *fiber.Wrapper) *server {
	return &server{
		app,
	}
}
