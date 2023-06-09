package fiber

import (
	"fmt"

	"vm-controller/configs"
	"vm-controller/internal/core/infra/network/fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Wrapper struct {
	App     *fiber.App
	Address string
}

func New(cfg configs.FiberConfig) *Wrapper {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               cfg.Prefork,
		CaseSensitive:         cfg.CaseSensitive,
		StrictRouting:         cfg.StrictRouting,
		ServerHeader:          cfg.ServerHeader,
		AppName:               cfg.AppName,
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "accept,content-type",
		AllowCredentials: true,
		MaxAge:           1728000,
	}))
	app.Use(middleware.NewLogger())
	// app.Use(csrf.New())

	return &Wrapper{
		App:     app,
		Address: fmt.Sprintf(":%d", cfg.Port),
	}
}
