package fiber

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type server struct {
	app *fiber.App
}

func New(conf config.FiberConfig) *server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               conf.Prefork,
		CaseSensitive:         conf.CaseSensitive,
		StrictRouting:         conf.StrictRouting,
		ServerHeader:          conf.ServerHeader,
		AppName:               conf.AppName,
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "accept,content-type,authorization,x-sic-device-id",
		AllowCredentials: true,
		MaxAge:           1728000,
	}))
	app.Use(middleware.NewLogger())
	// app.Use(csrf.New())

	return &server{
		app: app,
	}
}
