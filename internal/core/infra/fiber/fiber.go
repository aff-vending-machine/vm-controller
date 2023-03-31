package fiber

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func New(conf config.FiberConfig) *fiber.App {
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
		AllowHeaders:     "accept,content-type,authorization",
		AllowCredentials: true,
		MaxAge:           1728000,
	}))
	app.Use(logger.New())

	return app
}
