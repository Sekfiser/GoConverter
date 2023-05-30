package app

import (
	"Gonverter/app/router"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string) {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Converter",
		AppName:       "GoConverter v0.2",
	})

	app.Route(
		"/",
		router.RegisterEndpoints,
		"converter.",
	)

	log.Fatal(app.Listen(":" + port))
}
